# uxv

This is a memorial project for Unix and line editor ed.

**What is the meaning of ux and v?**

uxv = Unix V = Unix 5 = Unix fifth part (at least from my personal perspective)(yes I know about Plan 9, Inferno, Limbo and 9P):

1. Ancient Unix from Bell labs
2. BSD
3. Linux
4. macOS & Android
5. my experiments with edv editor, shv shell and uxv OS

**Why I'm doing this?**

Just for programming fun and enjoyment. When I turned 50 it was around the same time as ~50 years since the genesis of Unix. I like both Unix's story and system itself so I'll try to make extremely small and simple new ed, sh and Unix just to see what will happen. I plan for edv not to be a full screen editor like vi nor even a single line editor as ed. Instead edv will be a single character editor, as strange as it sounds, but with, I hope, some special and interesting features. If that happens to work well then edv might evolve to single line and more. After edv is finished I'll try to make a shv shell and maybe even a uxv operating system. And who knows what else if my imagination and motivation propagates further.

Programming of edv started in ed on Solaris as some kind of distant enactment of Ken writing ancient Unix software on qed and Model 33 Teletype.

**Why Solaris and ed editor?**

Solaris is Unix and I see BSD and Solaris as advanced versions of ancient Unix. When Solaris boots one can see 'Loading unix...' message which is nice to see. So both this Readme.md file and Go code for edv were entirely written on Solaris vm and in ed in their first iterations. Not an easy task and I believe that this is the main reason why ex and vi later appeared.

**Why not BSD or something else?**

OpenIndiana Solaris (distribution of Illumos OS based on Sun Solaris Express) was the only free original Unix which worked sufficiently well in virtual machine. Oracle (Sun) 11-3 or 11-4 Solaris didn't work well in vm, free install images are not for development or production but just for testing so their licence is not suitable if my project evolves further. NetBSD was used only for testing in the first several iterations of edv but not for development because of the weirdly high constant CPU loads in vm.

**Why Go and not C, C++, Rust, Zig?**

I look at Go, through C, as a distant advanced version of ancient B. And Ken who made B is also one of the creators of Go. Today C is mainly for systems programming, which is not my primary interest, and Go has better prospects for other applications at least for me in backend (and web development).

After programming in ed further work will continue in vi(m) and/or Emacs as they are fundamental parts of computer history and editor evolution. After that work will maybe continue in Visual Studio Code. Also main programming will include Fedora Linux with testing on Solaris vm.

**License: GPLv2.0**
