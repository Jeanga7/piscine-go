preuve=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
echo $preuve 
cat interviews/interview-$preuve
echo $MAIN_SUSPECT