DAY="day$1"

cp base.go.template "${DAY}.go"
cp base_test.go.template "${DAY}_test.go"
touch ./data/${DAY}.txt