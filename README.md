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
4. `phptelegoram -token telegram.token.file -php entry.php -w 5`

## TODO

* Support timed task. (non-blocked)

## License

Most lines in `php_embed.c` is forked from [Xing Xing's GoEmPHP](https://github.com/mikespook/goemphp), see the original project page for detailed license information.

You can choose ANY VERSION of MIT or GPL for the codes written by me.

