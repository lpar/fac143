fac143
======

Written in the wake of Shellshock, this is a small utility to search for
executables which might be susceptible to compromise.

It's written in Go to quickly scan one or more directories of files. It reports
executables which are shell or bash scripts, and also checks for setuid/setgid
on them. 

I've only been writing Go for a couple of days, so please be gentle with your
style critiques.

meta@pobox.com.
GPL v3 or higher. Golang. The deeper you get, the sweeter the pain.

Links
=====

 * [CVE record](http://web.nvd.nist.gov/view/vuln/detail?vulnId=CVE-2014-6271).
 * [Everything you need to know about shellshock](http://www.troyhunt.com/2014/09/everything-you-need-to-know-about.html).
 * [In the wake of shellshock](http://lpar.ath0.com/2014/09/26/wake-shellshock/).
 * [How to convert the bash scripts you find to POSIX shell scripts](http://drj11.wordpress.com/2014/09/26/10-tips-for-turning-bash-scripts-into-portable-posix-scripts/).

