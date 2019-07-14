# Backups up to XML the MySQL data from the sites

DATE=`date --iso-8601=date`
FOLDER="/some/path/$DATE"

echo $FOLDER
exit 1

if [ ! -d "$FOLDER" ]; then
	echo "Creating folder $FOLDER"
  	mkdir $FOLDER
else
	echo "Reusing folder $FOLDER"
fi

echo "Backing up...."
mysql -D blogdb -e "select * from blogs;" --xml > $FOLDER/hc.xml

mysql -D cookingdb -e "select * from recipes;" --xml > $FOLDER/recipes.xml

mysql -D hkdb -e "select * from blogs;" --xml > $FOLDER/hk.xml