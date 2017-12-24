#!/bin/bash
#create Database
dbname=$(whiptail --inputbox "Input Your Database Name !" 8 78 --title "Database Name" 3>&1 1>&2 2>&3)
if [ $? = 0 ]; then

dbuser=$(whiptail --inputbox "Input Your Database User !" 8 78 --title "Database User" 3>&1 1>&2 2>&3)
if [ $? = 0 ]; then

dbpass=$(whiptail --passwordbox "Input Your Password Database !" 8 78 --title "Database Password" 3>&1 1>&2 2>&3)
if [ $? = 0 ]; then

dbdir=$(whiptail --inputbox "Input Your New Direktori CMS Sekolaku !" 8 78 --title "Direktori CMS" 3>&1 1>&2 2>&3)
if [ $? = 0 ]; then

#konek database
Q1="CREATE DATABASE IF NOT EXISTS $dbname;"
Q2="GRANT ALL ON *.* TO '$dbuser'@'localhost' IDENTIFIED BY '$dbpass';"
Q3="FLUSH PRIVILEGES;"
SQL="${Q1}${Q2}${Q3}"

# mysql key
>~/.my.cnf
echo "[mysql]" >> ~/.my.cnf
echo "user=$dbuser" >> ~/.my.cnf
echo "password=$dbpass" >> ~/.my.cnf
mysql -uroot -e "$SQL"

#INstall file CMS
cp cms-sekolahku-v2.1.0.zip /var/www/html/
cd /var/www/html/
unzip cms-sekolahku-v2.1.0.zip
mv bwa/ $dbdir
chmod -R 755 $dbdir
chown -R www-data:www-data $dbdir
rm cms-sekolahku-v2.1.0.zip

#konek database dgn cms
cd /var/www/html/$dbdir/application/config
perl -pi -e "s/root/$dbuser/g" database.php
perl -pi -e "s/sekolah/$dbname/g" database.php
perl -pi -e "s/nur/$dbpass/g" database.php

#Konek database dgn sql
echo "Import Database"
cd /var/www/html/$dbdir/
mysql -uroot $dbname < db_cms_sekolahku.sql

#GAuge procees :v 
{
	for ((i = 0 ; i <= 100 ; i+=1)); do
	    sleep 0.3
	    echo $i
	    done
} | whiptail --gauge "Installing CMS" 6 50 0

## Result
echo "Sekolahku Successfully Installed" >> ini.txt
echo "You Must login from ipserver/$dbdir/login" >> ini.txt
echo "Login : admin" >> ini.txt
echo "Pass : admin" >> ini.txt
whiptail --textbox ini.txt 12 80
rm ini.txt
rm ~/.my.cnf

else
echo "Abort Installation"
fi

else 
echo "Abort Installation"
fi

else
echo "Abort Installation"
fi

else
echo "Abort Installation"
fi

else
echo "Abort Installation"
fi

else
echo "Abort Installation"
fi
