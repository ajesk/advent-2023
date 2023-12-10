DAY="day$1"
SOURCE="${DAY}.go"
TEST="${DAY}_test.go"

cp -v base.go.template $SOURCE
cp -v base_test.go.template $TEST
sed -i "s/\[x\]/${DAY}/g" $SOURCE
sed -i "s/\[x\]/${DAY}/g" $TEST
touch ./data/${DAY}.txt
touch ./data/${DAY}demo.txt