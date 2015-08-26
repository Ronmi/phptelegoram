<?php

function processMessage(Message $message) {
    // reply to user message
    $opt = new SendOptions;
    $opt->replyTo($message);

    // reply after 60 seconds
    timed_task(60*1000, array(
        'uid' => $message->senderID(),
        'message' => '[Hello, ' . $message->text() . ']',
        'opt' => $opt->json(),
    ));
}

function processTask(array $data = null) {
    send_message($data['uid'], $data['message'], SendOptions::fromJson($data['opt']));
}
