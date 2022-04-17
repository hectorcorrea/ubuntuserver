# Run this via a cron job every X number of days
DATE=`date --iso-8601=date`
HC_FOLDER="/root/backups/$DATE/hc/"
COOKING_FOLDER="/root/backups/$DATE/cooking/"

echo "Creating folders for $DATE ..."
mkdir -p $HC_FOLDER
mkdir -p $COOKING_FOLDER

echo "Backing up on $DATE ..."
cp -r /root/hectorcorrea.com/data/ $HC_FOLDER
cp -r /root/cooking/data/ $COOKING_FOLDER
echo "Done."
