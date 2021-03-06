/*
 * l2_fib.h : layer 2 forwarding table (aka mac table)
 *
 * Copyright (c) 2013 Cisco and/or its affiliates.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#ifndef included_l2fib_h
#define included_l2fib_h

#include <vlib/vlib.h>
#include <vppinfra/bihash_8_8.h>

/*
 * The size of the hash table
 */
#define L2FIB_NUM_BUCKETS (256 * 1024)
#define L2FIB_MEMORY_SIZE (128<<20)

/* Ager scan interval is 1 minute for aging */
#define L2FIB_AGE_SCAN_INTERVAL		(60.0)

/* MAC event scan delay is 100 msec unless specified by MAC event client */
#define L2FIB_EVENT_SCAN_DELAY_DEFAULT	(0.1)

/* Max MACs in a event message is 100 unless specified by MAC event client */
#define L2FIB_EVENT_MAX_MACS_DEFAULT	(100)

/* MAC event learn limit is 1000 unless specified by MAC event client */
#define L2FIB_EVENT_LEARN_LIMIT_DEFAULT	(1000)

typedef struct
{

  /* hash table */
  BVT (clib_bihash) mac_table;

  /* number of buckets in the hash table */
  uword mac_table_n_buckets;

  /* hash table memory size */
  uword mac_table_memory_size;

  /* hash table initialized */
  u8 mac_table_initialized;

  /* last event or ager scan duration */
  f64 evt_scan_duration;
  f64 age_scan_duration;

  /* delay between event scans, default to 100 msec */
  f64 event_scan_delay;

  /* max macs in event message, default to 100 entries */
  u32 max_macs_in_event;

  /* convenience variables */
  vlib_main_t *vlib_main;
  vnet_main_t *vnet_main;
} l2fib_main_t;

extern l2fib_main_t l2fib_main;

/*
 * The L2fib key is the mac address and bridge domain ID
 */
typedef struct
{
  union
  {
    struct
    {
      u16 bd_index;
      u8 mac[6];
    } fields;
    struct
    {
      u32 w0;
      u32 w1;
    } words;
    u64 raw;
  };
} l2fib_entry_key_t;

STATIC_ASSERT_SIZEOF (l2fib_entry_key_t, 8);

/**
 * A combined representation of the sequence number associated
 * with the interface and the BD.
 * The BD is in higher bits, the interface in the lower bits, but
 * the order is not important.
 *
 * It's convenient to represent this as an union of two u8s,
 * but then in the DP one is forced to do short writes, followed
 * by long reads, which is a sure thing for a stall
 */
typedef u16 l2fib_seq_num_t;

static_always_inline l2fib_seq_num_t
l2_fib_mk_seq_num (u8 bd_sn, u8 if_sn)
{
  return (((u16) bd_sn) << 8) | if_sn;
}

static_always_inline l2fib_seq_num_t
l2_fib_update_seq_num (l2fib_seq_num_t sn, u8 if_sn)
{
  sn &= 0xff00;
  sn |= if_sn;

  return (sn);
}

extern void l2_fib_extract_seq_num (l2fib_seq_num_t sn, u8 * bd_sn,
				    u8 * if_sn);
extern u8 *format_l2_fib_seq_num (u8 * s, va_list * a);

/**
 * Flags associated with an L2 Fib Entry
 *   - static mac, no MAC move
 *   - not subject to age
 *   - mac is for a bridged virtual interface
 *   - drop packets to/from this mac
 *   - MAC learned to be sent in L2 MAC event
 *   -MAC learned is a MAC move
 */
#define foreach_l2fib_entry_result_attr       \
  _(STATIC,  0, "static")                     \
  _(AGE_NOT, 1, "age-not")                    \
  _(BVI,     2, "bvi")                        \
  _(FILTER,  3, "filter")                     \
  _(LRN_EVT, 4, "learn-event")                \
  _(LRN_MOV, 5, "learn-move")

typedef enum l2fib_entry_result_flags_t_
{
  L2FIB_ENTRY_RESULT_FLAG_NONE = 0,
#define _(a,v,s) L2FIB_ENTRY_RESULT_FLAG_##a = (1 << v),
  foreach_l2fib_entry_result_attr
#undef _
} __attribute__ ((packed)) l2fib_entry_result_flags_t;

STATIC_ASSERT_SIZEOF (l2fib_entry_result_flags_t, 1);

extern u8 *format_l2fib_entry_result_flags (u8 * s, va_list * args);

/*
 * The l2fib entry results
 */
typedef struct l2fib_entry_result_t_
{
  union
  {
    struct
    {
      u32 sw_if_index;		/* output sw_if_index (L3 intf if bvi==1) */
      l2fib_entry_result_flags_t flags;

      u8 timestamp;		/* timestamp for aging */
      l2fib_seq_num_t sn;	/* bd/int seq num */
    } fields;
    u64 raw;
  };
} l2fib_entry_result_t;

STATIC_ASSERT_SIZEOF (l2fib_entry_result_t, 8);

#define _(a,v,s)                                                        \
  always_inline int                                                     \
  l2fib_entry_result_is_set_##a (const l2fib_entry_result_t *r) {       \
    return (r->fields.flags & L2FIB_ENTRY_RESULT_FLAG_##a);             \
  }
foreach_l2fib_entry_result_attr
#undef _
#define _(a,v,s)                                                        \
  always_inline void                                                    \
  l2fib_entry_result_set_##a (l2fib_entry_result_t *r) {       \
    r->fields.flags |= L2FIB_ENTRY_RESULT_FLAG_##a;             \
  }
  foreach_l2fib_entry_result_attr
#undef _
#define _(a,v,s)                                                        \
  always_inline void                                                    \
  l2fib_entry_result_clear_##a (l2fib_entry_result_t *r) {       \
    r->fields.flags &= ~L2FIB_ENTRY_RESULT_FLAG_##a;             \
  }
  foreach_l2fib_entry_result_attr
#undef _
  static inline void
l2fib_entry_result_set_bits (l2fib_entry_result_t * r,
			     l2fib_entry_result_flags_t bits)
{
  r->fields.flags |= bits;
}

static inline void
l2fib_entry_result_clear_bits (l2fib_entry_result_t * r,
			       l2fib_entry_result_flags_t bits)
{
  r->fields.flags &= ~bits;
}

/* L2 MAC event entry action enums (see mac_entry definition in l2.api) */
typedef enum
{
  MAC_EVENT_ACTION_ADD = 0,
  MAC_EVENT_ACTION_DELETE = 1,
  MAC_EVENT_ACTION_MOVE = 2,
} l2_mac_event_action_t;

/**
 * Compute the hash for the given key and return
 * the corresponding bucket index
 */
always_inline u32
l2fib_compute_hash_bucket (l2fib_entry_key_t * key)
{
  u32 result;
  u32 temp_a;
  u32 temp_b;

  result = 0xa5a5a5a5;		/* some seed */
  temp_a = key->words.w0;
  temp_b = key->words.w1;
  hash_mix32 (temp_a, temp_b, result);

  return result % L2FIB_NUM_BUCKETS;
}

always_inline u64
l2fib_make_key (const u8 * mac_address, u16 bd_index)
{
  u64 temp;

  /*
   * The mac address in memory is A:B:C:D:E:F
   * The bd id in register is H:L
   */
#if CLIB_ARCH_IS_LITTLE_ENDIAN
  /*
   * Create the in-register key as F:E:D:C:B:A:H:L
   * In memory the key is L:H:A:B:C:D:E:F
   */
  temp = CLIB_MEM_OVERFLOW_LOAD ((u64 *) mac_address) << 16;
  temp = (temp & ~0xffff) | (u64) (bd_index);
#else
  /*
   * Create the in-register key as H:L:A:B:C:D:E:F
   * In memory the key is H:L:A:B:C:D:E:F
   */
  temp = CLIB_MEM_OVERFLOW_LOAD ((u64 *) mac_address) >> 16;
  temp = temp | (((u64) bd_index) << 48);
#endif

  return temp;
}



/**
 * Lookup the entry for mac and bd_index in the mac table for 1 packet.
 * Cached_key and cached_result are used as a one-entry cache.
 * The function reads and updates them as needed.
 *
 * mac0 and bd_index0 are the keys. The entry is written to result0.
 * If the entry was not found, result0 is set to ~0.
 *
 * key0 return with the computed key, convenient if the entry needs,
 * to be updated afterward.
 */

static_always_inline void
l2fib_lookup_1 (BVT (clib_bihash) * mac_table,
		l2fib_entry_key_t * cached_key,
		l2fib_entry_result_t * cached_result,
		u8 * mac0,
		u16 bd_index0,
		l2fib_entry_key_t * key0, l2fib_entry_result_t * result0)
{
  /* set up key */
  key0->raw = l2fib_make_key (mac0, bd_index0);

  if (key0->raw == cached_key->raw)
    {
      /* Hit in the one-entry cache */
      result0->raw = cached_result->raw;
    }
  else
    {
      /* Do a regular mac table lookup */
      BVT (clib_bihash_kv) kv;

      kv.key = key0->raw;
      kv.value = ~0ULL;
      BV (clib_bihash_search_inline) (mac_table, &kv);
      result0->raw = kv.value;

      /* Update one-entry cache */
      cached_key->raw = key0->raw;
      cached_result->raw = result0->raw;
    }
}


/**
 * Lookup the entry for mac and bd_index in the mac table for 2 packets.
 * The lookups for the two packets are interleaved.
 *
 * Cached_key and cached_result are used as a one-entry cache.
 * The function reads and updates them as needed.
 *
 * mac0 and bd_index0 are the keys. The entry is written to result0.
 * If the entry was not found, result0 is set to ~0. The same
 * holds for mac1/bd_index1/result1.
 */
static_always_inline void
l2fib_lookup_2 (BVT (clib_bihash) * mac_table,
		l2fib_entry_key_t * cached_key,
		l2fib_entry_result_t * cached_result,
		u8 * mac0,
		u8 * mac1,
		u16 bd_index0,
		u16 bd_index1,
		l2fib_entry_key_t * key0,
		l2fib_entry_key_t * key1,
		l2fib_entry_result_t * result0,
		l2fib_entry_result_t * result1)
{
  /* set up key */
  key0->raw = l2fib_make_key (mac0, bd_index0);
  key1->raw = l2fib_make_key (mac1, bd_index1);

  if ((key0->raw == cached_key->raw) && (key1->raw == cached_key->raw))
    {
      /* Both hit in the one-entry cache */
      result0->raw = cached_result->raw;
      result1->raw = cached_result->raw;
    }
  else
    {
      BVT (clib_bihash_kv) kv0, kv1;

      /*
       * Do a regular mac table lookup
       * Interleave lookups for packet 0 and packet 1
       */
      kv0.key = key0->raw;
      kv1.key = key1->raw;
      kv0.value = ~0ULL;
      kv1.value = ~0ULL;

      BV (clib_bihash_search_inline) (mac_table, &kv0);
      BV (clib_bihash_search_inline) (mac_table, &kv1);

      result0->raw = kv0.value;
      result1->raw = kv1.value;

      /* Update one-entry cache */
      cached_key->raw = key1->raw;
      cached_result->raw = result1->raw;
    }
}

static_always_inline void
l2fib_lookup_4 (BVT (clib_bihash) * mac_table,
		l2fib_entry_key_t * cached_key,
		l2fib_entry_result_t * cached_result,
		const u8 * mac0,
		const u8 * mac1,
		const u8 * mac2,
		const u8 * mac3,
		u16 bd_index0,
		u16 bd_index1,
		u16 bd_index2,
		u16 bd_index3,
		l2fib_entry_key_t * key0,
		l2fib_entry_key_t * key1,
		l2fib_entry_key_t * key2,
		l2fib_entry_key_t * key3,
		l2fib_entry_result_t * result0,
		l2fib_entry_result_t * result1,
		l2fib_entry_result_t * result2,
		l2fib_entry_result_t * result3)
{
  /* set up key */
  key0->raw = l2fib_make_key (mac0, bd_index0);
  key1->raw = l2fib_make_key (mac1, bd_index1);
  key2->raw = l2fib_make_key (mac2, bd_index2);
  key3->raw = l2fib_make_key (mac3, bd_index3);

  if ((key0->raw == cached_key->raw) && (key1->raw == cached_key->raw) &&
      (key2->raw == cached_key->raw) && (key3->raw == cached_key->raw))
    {
      /* Both hit in the one-entry cache */
      result0->raw = cached_result->raw;
      result1->raw = cached_result->raw;
      result2->raw = cached_result->raw;
      result3->raw = cached_result->raw;
    }
  else
    {
      BVT (clib_bihash_kv) kv0, kv1, kv2, kv3;

      /*
       * Do a regular mac table lookup
       * Interleave lookups for packet 0 and packet 1
       */
      kv0.key = key0->raw;
      kv1.key = key1->raw;
      kv2.key = key2->raw;
      kv3.key = key3->raw;
      kv0.value = ~0ULL;
      kv1.value = ~0ULL;
      kv2.value = ~0ULL;
      kv3.value = ~0ULL;

      BV (clib_bihash_search_inline) (mac_table, &kv0);
      BV (clib_bihash_search_inline) (mac_table, &kv1);
      BV (clib_bihash_search_inline) (mac_table, &kv2);
      BV (clib_bihash_search_inline) (mac_table, &kv3);

      result0->raw = kv0.value;
      result1->raw = kv1.value;
      result2->raw = kv2.value;
      result3->raw = kv3.value;

      /* Update one-entry cache */
      cached_key->raw = key1->raw;
      cached_result->raw = result1->raw;
    }
}

void l2fib_clear_table (void);

void l2fib_table_init (void);

void
l2fib_add_entry (const u8 * mac,
		 u32 bd_index,
		 u32 sw_if_index, l2fib_entry_result_flags_t flags);

static inline void
l2fib_add_filter_entry (const u8 * mac, u32 bd_index)
{
  l2fib_add_entry (mac, bd_index, ~0,
		   (L2FIB_ENTRY_RESULT_FLAG_FILTER |
		    L2FIB_ENTRY_RESULT_FLAG_STATIC));
}

u32 l2fib_del_entry (const u8 * mac, u32 bd_index, u32 sw_if_index);

void l2fib_start_ager_scan (vlib_main_t * vm);

void l2fib_flush_int_mac (vlib_main_t * vm, u32 sw_if_index);

void l2fib_flush_bd_mac (vlib_main_t * vm, u32 bd_index);

void l2fib_flush_all_mac (vlib_main_t * vm);

void
l2fib_table_dump (u32 bd_index, l2fib_entry_key_t ** l2fe_key,
		  l2fib_entry_result_t ** l2fe_res);

u8 *format_vnet_sw_if_index_name_with_NA (u8 * s, va_list * args);

BVT (clib_bihash) * get_mac_table (void);

#endif

/*
 * fd.io coding-style-patch-verification: ON
 *
 * Local Variables:
 * eval: (c-set-style "gnu")
 * End:
 */
