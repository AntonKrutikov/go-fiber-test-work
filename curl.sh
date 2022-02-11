curl 'http://localhost:3000/echo?id=1&name=Test'
echo
curl --data 'id=1&name=Test' "http://localhost:3000/echo"
echo
curl --data 'name=Test' "http://localhost:3000/echo?id=1"
echo
curl -X POST -H "Content-Type: application/json" http://localhost:3000/echo -d '{"id":1,"name":"Test"}'
echo
curl 'http://localhost:3000/echo?id=string_instead_int&name=Test'
echo
curl 'http://localhost:3000/echo'