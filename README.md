# imdb-chart-fetcher
## How to run the assignment

1. Update all the dependencies
   <pre>go get -u</pre>
2. Run command 
   <pre>go build</pre>
3. Run command as per OS. I am using a windows machine to run my assignment 
   <pre>imdb_chart_fetcher 'https://www.imdb.com/india/top-rated-indian-movies' 1 </pre>
   
### OR if you are running the assignment on an Windows machine 
1. Open command line window in the imdb_chart_fetcher directory
2. Run 
<pre>imdb_chart_fetcher 'https://www.imdb.com/india/top-rated-indian-movies' 1 </pre>


Example:
<pre>
Input --
>imdb_chart_fetcher 'https://www.imdb.com/india/top-rated-indian-movies' 1

Output -- 

[{"title":"Pather Panchali","movie_release_year":1955,"imdb_rating":8.6,"summary":"Impoverished priest Harihar Ray, dreaming of a better life for himself and his family, leaves his rural Bengal village in search of work.","duration":"2h 5min","genre":"Drama"}]
</pre>
   
##### The repository contains a build based on my machine configuration i.e. Windows 10 x64 arch

