set datafile separator " "
set grid
set autoscale 
set output "filter.png"
set terminal png size 1024, 768

f = "/tmp/filter.txt"

plot f u 1:2 w points lc rgb 'red' title 'X', \
    f u 1:3 w lines lc rgb 'blue' title 'Filtered', \
    f u 1:4 w lines lc rgb 'black' title 'SMA(benchmark)'

