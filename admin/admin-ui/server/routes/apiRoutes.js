var express = require('express');
var router = express.Router();
const axios = require('axios');
const https = require('https');
const Settings = require('../../client/public/conf/Settings.js');

/* GET users listing. */
router.get('/', async function (req, res, next) {
    const accessToken = req.cookies.access_token;

    const instance = axios.create({
        httpsAgent: new https.Agent({
            rejectUnauthorized: false
        })
    });
    const headers = {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${accessToken}`,
    };
    console.log("*****************");
    console.log(headers);
    console.log("*****************");
    try {
        const response = await instance.get(`${Settings.app.rest_api}/application-rate-plans`, {
            headers: headers
        });
        const data = response.data;
        console.log(data);
        res.status(200).json(data);
    } catch (error) {
        console.log(error);
        res.status(500).json({ error: error.message });
    }


    // Retrieve the access token from the HTTP-only cookie

    // Make a request to the API server with the access token
    // Assuming the API server is located at https://api.example.com/items

});


module.exports = router;
