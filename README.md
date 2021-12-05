# Overview
[![Twitter](https://img.shields.io/badge/author-%40MachielMolenaar-blue.svg)](https://twitter.com/MachielMolenaar)

This is the source of StrangerBot, the bot on Telegram that matches two random
users and allows them to chat with each other.

# License
StrangerBot is licensed under the Apache 2.0 license.

# Installation
Currently there are no binaries available as direct download yet, you should
build it yourself using Go.

If you have go installed, you can install strangerbot like this:

`go get -u github.com/Machiel/strangerbot`

# How Run?

1. cd /root/go/src/strangerbot dir
2. run 'go install .'
3. cd /root/go/bin
4. modify start.sh,edit telegram bot key
5. sh start.sh

# Usage

Make sure you have MySQL installed, and retrieved an API key from Telegram.

## Example

Make sure you have the following environment variables set:

```
MYSQL_USER
MYSQL_PASSWORD
MYSQL_DATABASE
TELEGRAM_BOT_KEY
```

You can then run start StrangerBot by running `strangerbot` in your terminal.

## Menu configuration

| Field | Default | Description |
| --- | --- | --- |
| parent_id | 0 | parent menu id 0.top menu >0.sub menu |
| target_type | 1 | target 1.menu to menu 2. menu to question |
| name | | menu name(title) |
| question_id | | if target_type is 2, this field must be set question id |
| sort | 0 | order by sort asc |
| row_index | 0 | same row index will be inline |
| helper_title | | menu helper title |
| helper_text | | menu helper text |
| is_back_enabled | 1 | show back button 0.false 1.true |
| back_button_text | Back | back button text |

## Question configuration

| Field | Default | Description |
| --- | --- | --- |
| scene_type | 0 | scene type 1.profile question 2.matching question |
| helper_title | | question helper title |
| title | | question title |
| helper_text | | helper text |
| frontend_type | 1 | frontend type 1.select 2.multi select |
| max_multi_len | 0 | multi select max choose option length, 0 is not limits. |
| sort | 0 | order by sort asc, this field not used. |
| matching_mode | 1 | order by sort asc, this field not used. |
| matching_question_id | 0 | if scene is matching , it is origin question id, if 0 it will not support matching |

## Option configuration

| Field | Default | Description |
| --- | --- | --- |
| question_id | 0 | question id it form form_question table |
| option_type | 1 | option type 1.value option, now only have one type. |
| matching_option_id | 0 | match option id |
| label | | option label. |
| value | | option value. |
| is_matching_any | 0 | 0.false 1.true |
| sort | 0 | order by sort asc |
| row_index | 0 | same row index will be inline |


