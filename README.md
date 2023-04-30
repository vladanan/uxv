# uxv

uxv is project to learn unix and go.

uxv = Unix V = Unix 5 = Unix fifth part (at least for me):

1. Unix from Bell labs
2. BSD
3. Linux
4. MacOs & Android
5. my experiment with uxv, ed and sh

**Why edv and shv?**

The same reason as for uxv.

**Why this?**

It is when I'm 50 I met and learned about Unix and it is also ~50 years from genesis of Unix.
And i like Unix story and system so i'll try to make new ed and sh just for fun and enjoyment.
I plan to write my edv on ed, just like Ken wrote first Unix software on qed and on Model 33 Teletype.
Later if things get complex I'll continue with vi (and afther that with Visual studio code or something similar).

**Why NetBSD?**

Only BSD which is working good enough in VM, while I'm stll trying to make OI Solaris work at satisfactory level.

**Why (Open indiana) Solaris?**

The only free original Unix which worked well in virtual machine.
Oracle (Sun) 11-3 or 11-4 Solaris didn't work well in vm, install images are not for dev or prod or commerce
but just for testing so their licence is not suitable if my project evolves further.
But OI is not fully functional because of problem with scroll for mouse, so BSD was be the first option until scroll is fixed.
Later I found out that in vm BSD runs cpu at around 80% for no apparent reason so I switched back to IO Solaris because
scroll problem is not big (and can be fixed by little hacking in vm setup).

uxv development on Open indiana (based on Sun Solaris express) and tests on Fedora, Oracle Solaris, NetBSD.

**Why virtual machine?**

Because I don't have another comp and even if I had I have no more space on my desk for it.

**Why go and not c, c++, rust etc.?**

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

Some notes while editing this readme file with ed. As far as I know there is no easy way to split line in two parts
and also to add characters or words at the start or end of line. While thinking about this I realised that when
working on assembler code there is probablu no need for that kind of editing. I have found on some site the original
asm code for first Unix and simple operations in ed would be probably just enough to work on that bur for editing
text like this one and to split and move parts of lines is not easy in ed. So ed is probably just fine and good for
asm code and it was not intended to be used with readme files probably.

The same impressions I have when working with Go code. Changing part of line to print other text involves either completely
new line or s command whith more or less work. So it seems that ed was also not meant for higher level languages and even
if it was used for that it was pretty tough, especcialy on Teletypes 50 years ago. So, maybe, I'll make switch to vi sooner
than I expected XD

**License: GPLv2.0**
