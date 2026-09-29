# Way of working

This repository is read-only, fork it and work on your copy just as you would on a live one. \
Once you are satisfied with your work, please send us your repository url.

# Overview

This sample repo is pretty close to what we do at Leboncoin. \
The scope you own is found under `/ads`. \
You can also modify what is found under `/common`. \
Do not change any code under `/users`.

# Instructions

Here is the list of things to do:
1) There is a bug in the `GetAdByID` usecase, the usecase does not work, fix it.

2) Add a delete ad by id endpoint, usecase and dao method.

3) In the `GetAdByID` endpoint, add the username field based on the information found in the user service. Use the provided client in `users/pkg/client.go`.

4) Users are reporting that the `list_ads` endpoint is very slow when using the `user_id` filter, implement a fix.

5) Users are reporting that the bulk import is not working properly, it fails sometimes and when they try to run again, some ads appear twice, implement a fix.

The complexity is more or less increasing with each item, do not beat yourself up if you miss some things.

Finally, feel free to improve the code in any way that seem appropriate to you or fix any latent issue: the code you push should reflect what you feel comfortable shipping in production.
