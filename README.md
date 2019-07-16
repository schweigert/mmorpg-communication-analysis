# MMORPG Communication Analysis
This repository is part of my bachelor thesis.
I implemented Rudy (Tibia), Salz (Albion) and Willson (GuildWars 1/2) microsservices architectures for a communication analisys using Golang. You can use for some tests :)

# Domain

The domain implements business rule methods.
If you add new business rules, you need to consider adding this new method to the services.

The rules are:
 - An `Account` can be created from a user and password;
 - A `Character` is created from an account, a name is required.
 - When you create a character, you are added to the center of a `Scene`.
 - A scene stores information adds to the game environment. It will concentrate generic information for the game service.