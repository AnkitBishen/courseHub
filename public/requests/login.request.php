<?php


$phpInp = file_get_contents('php://input');
$postData = json_decode($phpInp);


if (isset($postData->authentication) && $postData->authentication == true) {
    $email = $postData->email;
    $password = $postData->password;
}