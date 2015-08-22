<?php

$data = json_decode($message, true);
$sender = $data['from']['id'];

        
send_message($sender, "[Hello, " . $data['text'] . "]");