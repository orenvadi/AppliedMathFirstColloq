for i in {1..10}; do                                                                                                                ─╯
  echo "Running program $i..."
  go run $i.go > output_$i.txt 2>&1
  echo "Program $i completed. Output saved to output_$i.txt"
done

