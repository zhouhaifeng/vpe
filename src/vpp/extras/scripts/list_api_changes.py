#!/usr/bin/env python3
from __future__ import print_function
import fnmatch
import os
import subprocess

starttag = 'v19.08-rc0'
endtag = 'HEAD'
emit_md = True
apifiles = []

for root, dirnames, filenames in os.walk('.'):
    for filename in fnmatch.filter(filenames, '*.api'):
        apifiles.append(os.path.join(root, filename))

for f in apifiles:
    commits = subprocess.check_output(['git', 'log',
                                       '--oneline', starttag + '..' + endtag,
                                       f])
    if commits:
        if f[0:2] == './':
            f = f[2:]
        if emit_md:
            print("| @c %s ||" % f)
            print("| ------- | ------- |")
            for line in commits.splitlines():
                parts = line.strip().split()
                commit = parts[0]
                message = b" ".join(parts[1:]).decode().replace("|", r"\|")
                print("| [%s](https://gerrit.fd.io/r/gitweb?"
                      "p=vpp.git;a=commit;h=%s) | %s |" % (
                            commit, commit, message))
            print()
        else:
            print(f)
            print(commits)
