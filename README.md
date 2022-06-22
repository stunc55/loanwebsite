What is loanwebsite

This website was created so that students can borrow products through the website. So basically it is created for schools and other institutions so they can barrow products for students. The website what I have created is very simple and basic, but it is functional.  

Getting started

Welcome, in this readme file you can read everything you need to know about this code. I will share here important information before you can start using all files.

Name

The name of my code/product is called loanwebsite. Its just an basic name and it’s just created for educational purpose. So feel free to edit or change my code to use it. 

Version
This is the first version.
V1.0

Description

Like I said up here this code is created for educational purpose. This is also the first time I have uploaded files on github. I have worked on a project for school. For this project I had to create an website in Golang.  

Note!: The code includes Dutch designation. I will translate it in this readme file.

Functions

you can use the following functions on the website:
- Zoek producten (search products)
- Producten toevoegen (add products)
- Producten verwijdern (delete products)
- Producten wijzigen (edit products)

Features

Expections for: 2.0V
- Create login users
- Give grant permissions for users that are hosting the website
- Create pictures of the product
- Make the code more robust
If you have more ideas don’t hesitate to share it with me.

How to use

First of all, if you want to make changes to this code. Note that you have to change everything in the code including MySQL. 

Second of all, because it was made for students they don’t need to follow this description below.
They will have an different connection to the website with LDAP and VPN. This is and won’t be shared.

Check the following before using:
- Installed Docker (See version in code)
- Installed Portainer (Recommended) (Latest Version)
- Installed vscode (See version in code)
 	- installed go (See version in code)

If everything is installed you can continue reading how to use. 

You have to copy all the files in your Host (Linux) and open it in vscode.

Check of there isn’t any problem and if there is try to fix it. Like a small problem that a package isn’t installed. 
If everything looks fine, save all the files and remember where you located it on your desktop.

Open your Terminal and write (copy) the following command:
cd example/example
(write instead of example the location where you copied all the files)

Now you have to use the following command:
sudo docker-compose up -d

Now it will install all containers:
- MySQL
- phpMyAdmin
- Loanwebsite

And you are able to use the website if you write (copy) this on your url:
localhost:8095

Project status

This is the end of my readmi file. While I have some features, this project is stopped completely.
I hope you can enjoy.

