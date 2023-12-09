DAY="day$1"

cp -v base.go.template "${DAY}.go"
cp -v base_test.go.template "${DAY}_test.go"
touch ./data/${DAY}.txt