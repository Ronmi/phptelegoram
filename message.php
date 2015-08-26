<?php

require("lib.php");
require($entry_file);
if (function_exists("processMessage")) {
    processMessage(new Message($message_json));
}
