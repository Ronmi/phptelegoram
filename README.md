# PHPTeleGOram

PHP Telegram Bot powered by Golang and libphp-embed.

By using php-embbed, there would be some problem with some extensions. You can visit [PHP official site](https://php.net/) for further information.

## Build requirements

* Golang 1.4+
* libphp-embed

## Synopsis

1. Download the source.
2. Build executable. (tested with 1.4 and 1.5)
3. Write Telegram message handler with PHP.
4. Run following command within source folder: `phptelegoram -token your.telegram.token.file -php your.app.entry.php -w 5`

## Go part

Go part will use long-polling method to fetch telegram message updates, which is powered by a great project [telebot](https://github.com/tucnak/telebot). After receiving message, it will execute [message.php](https://github.com/Ronmi/phptelegoram/blob/master/message.php), which will include [lib.php](https://github.com/Ronmi/phptelegoram/blob/master/lib.php) and execute the program entry.

By calling `timed_task($milliseconds, $data)` in your php script, Go part will spawn a goroutine, sleep for a while, and execute [timedtask.php](https://github.com/Ronmi/phptelegoram/blob/master/timedtask.php).

As the result, you have to run main program in the same folder containing `message.php`, `timedtask.php` and `lib.php`, or copy these files to your working directory.

## PHP part

[lib.php](https://github.com/Ronmi/phptelegoram/blob/master/lib.php) provides some convenient method to access or send message. Briefly:

1. class `Message` provides some shorthand mthod to access received message.
2. class `SendOptions` for some additional options when sending message, like custom keyboard, message replying, etc.
3. function `send_message` for sending message to specified user or chatroom.
4. function `timed_task` for sending event after few miliseconds.

You have to implement two global function as program entrance:

1. `function processMessage(Message $message)` is called after received message from telegram.
2. `function processTask(array data = null)` is called as the result of `timed_task` API.

See [example.php](https://github.com/Ronmi/phptelegoram/blob/master/example.php) as a simple example for using them.

## Walk further
You can write your own version of `message.php` and/or `timedtask.php`.

* `message.php` receives original message as json-encoded string in global variable `$message_json`.
* `timedtask.php` receives original data as json-encoded string in global variable `$data`.

There are two functions exported from Go part to help you do the job:

1. `raw_send_message(int, string, string)` sends telegram message to user.
2. `raw_timed_task(int, string)` sets up a timed task.

These functions are wrapped in [lib.php](https://github.com/Ronmi/phptelegoram/blob/master/lib.php).

## License

Most lines in `php_embed.c` and `php.go` is forked from [Xing Xing's GoEmPHP](https://github.com/mikespook/goemphp), see the original project page for detailed license information.

You can choose ANY VERSION of MIT or GPL for the codes written by me.
