package vars

var (
	// debug or prod
	RUN_MODE = "prod"
)

const (
	StartMatchingMessage    = `Looking for another student to match you with... Hold on! (This may take a while! Keep your notifications on!) **NOTE: If you send anything illegal here, your data will be handed over to the police. Your User ID is anonymous only until you break the rules. A police report for harassment/defamation will be filed if you pass off another user's contact as if it is yours.** To report a user, enter **/report (followed by a reason; don't leave blank)** into the chat. If chat with a user you want to report has already ended, then **do not** start a new chat—immediately contact the admin @aaldentnay . **Also, a note to some guys here—pls stop being thirsty on here because that scares new users away; I'm taking a huge leapt of faith when I set up a platform like this. Those reported will be PERMANENTLY banned.** Misuse of /report , if not accidental, can also result in ban.`
	NotProfileFinishMessage = "please /setup configure your profile"
	TopMenuMessage          = `Set up your profile with your gender, preferred matching gender, and interest here! You'll need those before you can use /start to join the queue.`
	HelpMessage             = `Help:

Use /start to start looking for a conversational partner, once you're matched you can use /end to end the conversation.

Use /report to report a user, use it as follows:
/report <reason>

Use /nopics to disable receiving photos, and /nopics if you want to enable it again.

HEAD OVER to @unichatbotchannel for rules, updates, announcements or info on how you can support the bot!

Sending images and videos are a beta functionality, but appear to be working fine.

If you require any help, feel free to contact @aaldentnay !`
	MatchedMessage           = `You have been matched, have fun! Your match is %s and is open to %s!`
	QuestionCallbackTest     = `This question can only have %d options at most`
	BanMessage               = `You are banned until %s`
	RegTipMessage            = `You’re done setting up! As a final step before we find you a match, we will need to know your school email address. Don’t worry, this will be anonymous to other users! This is just to verify you’re a student :)`
	NeedInputEmailMessage    = `What is your student email address?`
	SendEmailCodeMessage     = `We've e-mailed you a code. Please check your e-mail and Enter the code here to complete the verification.`
	OTPNoExistsMessage       = `Code is not exists,Please re-enter email or code.`
	OTPFailMessage           = `Code is not match,Please re-enter email or code.`
	OTPSuccessMessage        = `Register Success! Please click /start to match.`
	NotAccessRegisterMessage = `Sorry, you do not have the authorization to register.`
	RetryInputEmailMessage   = `Mail service send failed, Please re-enter email.`
)

// emoji UTF-8 from https://apps.timwhitlock.info/emoji/tables/unicode
const (
	ChooseMark = "\xE2\x9C\x85"
)

// set goals question id, it work for MatchedMessage.
const (
	GoalsQuestionId int64 = 7
)
