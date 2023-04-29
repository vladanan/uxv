# uxv

uxv is project to learn unix and go.

uxv = Unix V = Unix 5 = Unix fifth part (at least for me):

1. Unix from Bell labs
2. BSD
3. Linux
4. MacOs & Android
5. my experiment with uxv, ed and sh

Why edv and shv?

The same reason as for uxv.

Why this?

It is when I'm 50 I met and learned about Unix and it is also ~50 years from genesis of Unix.
And i like Unix story and system so i'll try to make new ed and sh just for fun and enjoyment.
I plan to write my edv on ed, just like Ken wrote first Unix software on qed and on Model 33 Teletype.
Later if things get complex I'll continue with vi (and afther that with Visual studio code or something similar).

Why NetBSD?

Only BSD which is working good enough in VM, while I'm stll trying to make OI Solaris work at satisfactory level.

Why (Open indiana) Solaris?

The only free original Unix which worked well in virtual machine.
Oracle (Sun) 11-3 or 11-4 Solaris didn't work well in vm, install images are not for dev or prod or commerce
but just for testing so their licence is not suitable if my project evolves further.
But OI is not fully functional because of problem with scroll for mouse, so BSD was be the first option until scroll is fixed.
Later I found out that in vm BSD runs cpu at 70% for no apparent reason so I switched back to IO Solaris because
scroll problem is not big (and maybe can be temporarily fixed by little hacking in to vm setup).

uxv development on Open indiana (based on Sun Solaris express) and tests on Fedora, Oracle Solaris, NetBSD.

Why virtual machine?

Because I don't have another comp and even if I had I have no more space on my desk for it.

Why go and not c, c++, rust etc.?

BSD and Solaris are distant advanced versions of ancient Unix and
go through c is advanced version of b and
Ken who made b is also one of the creators of go and
today c is mainly for system programmers and go could have better prospect to learn
for other applications.

So first version would be written in ed
seccond in vi
third in sam?
fourth in emacs
fifth in VS Code.

**license: GPLv2.0**
