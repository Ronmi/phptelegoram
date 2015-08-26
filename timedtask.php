<?php

require("lib.php");
require($entry_file);
if (function_exists("processTask")) {
    processTask(json_decode($data, true));
}
