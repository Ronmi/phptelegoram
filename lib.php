<?php

class Message {
    private $data;

    public function __construct($json) {
        $this->data = json_decode($json, true);
    }

    public function get($key) {
        $arr = explode(".", $key);
        $ret = $this->data;
        foreach ($arr as $k => $v) {
            if (! isset($ret[$v])) {
                return null;
            }
            $ret = $ret[$v];
        }
        return $ret;
    }

    public function text() {
        return $this->get("text");
    }

    public function sender() {
        $this->get("from");
    }

    public function senderID() {
        return $this->get("from.id");
    }

    public function json() {
        return json_encode($this->data);
    }

    public static function fromJson($json) {
        return new self($json);
    }
}

class SendOptions {
    private $reply;
    private $markup;
    private $noPreview;

    public function __construct() {
        $this->markup = array("place_holder" => null);
        $this->noPreview = false;
    }

    public function replyTo(Message $msg) {
        $this->reply = $msg->json();
    }

    public function preview($bool) {
        $this->noPreview = !$bool;
    }

    public function forceReply($bool) {
        $this->markup['force_reply'] = $bool == true;
    }

    public function customKeyboard(array $keyboard) {
        $this->markup['keyboard'] = $keyboard;
    }

    public function resizeKeyboard($bool) {
        $this->markup['resize_keyboard'] = $bool == true;
    }

    public function onetime($bool) {
        $this->markup['one_time_keyboard'] = $bool == true;
    }

    public function hideKeyboard($bool) {
        $this->markup['hide_keyboard'] = $bool == true;
    }

    public function selective($boole) {
        $this->markup['selective'] = $bool == true;
    }

    public function json()
    {
        $data = array(
            "ReplyTo" => json_decode($this->reply, true),
            "ReplyMarkup" => $this->markup,
            "DisableWebPagePreview" => $this->noPreview,
        );

        $ret = json_encode($data);
        return $ret;
    }

    public static function fromJson($json) {
        $data = json_decode($json, true);
        $ret = new self;
        $ret->reply = json_encode($data['ReplyTo']);
        $ret->markup = $data['ReplyMarkup'];
        $ret->noPreview = $data['DisableWebPagePreview'];
        return $ret;
    }
}

function send_message($uid, $msg, SendOptions $options = null) {
    if ($options == null) {
        $options = new SendOptions;
    }
    raw_send_message($uid, $msg, $options->json());
}

function timed_task($milisec, $data) {
    raw_timed_task($milisec, json_encode($data));
}
