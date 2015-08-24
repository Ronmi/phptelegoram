<?php

$opt = new SendOptions;
$opt->replyTo($message);
send_message($message->senderID(), "[Hello, " . $message->text() . "]", $opt);
