const express = require('express')
const app = express()
const port = process.env.PORT || 3000;


app.use('/public', express.static(__dirname + '/public'))

app.get('/', (req, res) => {
  res.sendFile(__dirname + '/views/index.html')
})

app.get('/home', (req, res) => {
    res.sendFile(__dirname + '/views/home.html')
})

app.listen(port, () => {
  console.log('Server online')
})