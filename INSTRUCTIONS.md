# Way of working

This repository is read-ony, fork it and work on your copy just as you would on a live one.

Once you are satisfied with your work, please send us your repository url.

# Overview

This sample repo is pretty close to what we do at Leboncoin.

The scope you own is found under `/ads`.

You can also modify what is found under `/common`.

Do not change any code under `/users`.

# Instructions

There will be a number of instructions regarding this code.
Commit the fixes and contact us when you are done.

Some points are glaring / simple issues, other are trickier, do not beat yourself up if you miss some things.

1/ There is a bug in the `GetAdByID` usecase, the usecase does not work, fix it.

2/ Add a delete ad by id endpoint, usecase and dao method.

3/ In the `GetAdByID` endpoint, add the username field based on the information found in the user service. Use the provided client.

4/ Users are reporting that the `list_ads` endpoint is very slow when using the `user-id` filter, propose a fix.

5/ Users are reporting that the bulk import is not working properly, it fails sometimes and when they try to run again, some ads appear twice, propose a fix.

Finally, feel free to improve the code in any way that seem appropriate to you or fix any latent issue: the code you push should reflect what you feel comfortable shipping in production.
