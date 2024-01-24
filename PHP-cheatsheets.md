```php
<?php

// 退出进程并打印
die("This file is not ment to be ran. ¯\_(ツ)_/¯");
exit("This file is not ment to be ran. ¯\_(ツ)_/¯");

/**
 * Printing
 */
echo ""; // 打印字符串类型
print_r($arr); // Print anything, with type hints for array's and object's
var_dump($arr); // Print anything, with type hints for any value and sizes

/**
 * 字符串操作
 */
$string = 'Awesome cheatsheets';

str_contains($string, 'cheat'); // 查找字符串 是否包含指定的 字符串
str_replace('Awesome', 'Bonjour', $string); // 替换字符串
strcmp($string, 'Awesome cheatsheets'); // 比较字符串大小
strpos($string, 'a', 0); // 查找字符串首次出现的位置
str_split($string, 2); // 拆分字符串成数组
strrev($string); // 翻转字符串
trim($string); // 去除字符串首尾处的空白字符（或者其他字符）
ucfirst($string); // 使字符串的第 一个字符大写 
lcfirst($string); // 使字符串的第 一个字符小写
substr($string, 0, 4); // 返回子串

/**
 * Sorting an Array
 */
sort($arr); // 值 升序 不保留原键
rsort($arr); // 值 降序 不保留原键
asort($arr); // 值 升序 保留原键
ksort($arr); // 键 升序 保留原键
arsort($arr); // 值 降序 保留原键
krsort($arr); // 键 降序 保留原键

/**
 * Ways of looping
 */
continue; // 跳出本次循环的剩余代码，开启下一次循环，后台加参数指定跳出的层级
break; // 跳出当前循环，并且可以指定跳出层级


// Switch
switch($arr) {
    case 1:
        break;
    case 2:
        break;
    case 3:
        break;
    default:
}

/**
 * Match (PHP >= 8.0)
 * https://www.php.net/manual/fr/control-structures.match.php
 match 表达式跟 switch 语句相似，但是有以下关键区别：
match 比较分支值，使用了严格比较 (===)， 而 switch 语句使用了松散比较。
match 表达式会返回一个值。
match 的分支不会像 switch 语句一样， 落空时执行下个 case。
match 表达式必须彻底列举所有情况。
 */

$food = 'apple';
$return_value = match($food) {
    'apple', 'appel' => 'An apple',
    'banana' => 'A banana',
    'applepie' => 'An applepie',
    default => 'A fruit'
};

//You can also use it as a conditionnal and throw exceptions
$str = 'Welcome to awesome cheatsheets';
$return_value = match(true) {
    str_contains($str, 'Welcome') && str_contains($str ,'to') => 'en-EN',
    str_contains($str, 'Bonjour') && str_contains($str, 'sur') => 'fr-FR',
    default => throw new Exception('Not a recognized language')
};

/**
 * Global variables
 * http://php.net/manual/en/language.variables.superglobals.php
 */
$_SERVER; // SERVER variables
$_GET; // Query params
$_POST; // Post fields
$_REQUEST; // GET and POST together
$GLOBALS; // Array of global variables 全局变量 数组
$_SESSION; // Browser session 浏览器 会话
$_FILES; // Array of files that are sent in request
$_COOKIE; // Array of cookies sent in request
$_ENV; // php.ini options
$argv; // Array of terminal arguments (filename included) 命令行执行文件时的参数
$argc; // Number of arguments passed into terminal

/**
 * Functions
 */

 // Simple function
 function name($parameter);

 // Function with return type (void, int, float, string, array, object, mixed)
 function name($parameter) : void;

 // Function with optionnal parameter
 function name($parameter = '') : string;

 // Function with typed parameter (? means "can be null") 
 function name(?string $parameter) : ?string;

 // Function with union types (PHP >= 8.0)
 function name(int|string $parameter1, array $parameter2) : int|string;

 // Function call
 name('my_parameter');

 // Null safe operator (PHP >= 8.0)
 $myObject?->getName()?->startWith('A');

/**
 * Class 
 * http://php.net/manual/en/language.oop5.basic.php
 */
class NormalClass extends AbstractClassName implements InterfaceName
{

    // 引用 Trait (为 PHP 提供多继承的能力，可理解为代码复制)
    use TraitName;

    // --> 类属性类型 <--

    /**
     * 公有的类成员可以在任何地方被访问，会被继承。
     * @var Type
     */
    public $property;

    /**
     * 私有的类成员则只能被其定义所在的类访问，不会被继承。
     * @var Type
     */
    private $property;

    /**
     * 受保护的类成员则可以被其自身以及其子类和父类访问，会被继承。
     * @var Type
     */
    protected $property;

    /**
     * 静态变量，也被称为类变量，所有对象的变量都是同一个。
     * @var Type
     */
    static $property;

    // --> 方法类型 <--

    /**
     * 公共方法，任何对象都能访问。
     * @param Type
     * @return Type
     */
    public function publicFunction(Type $var = null): Type
    {
    }

    /**
     * 私有方法，只有对象自身可以访问。
     * @param Type
     * @return Type
     */
    private function privateFunction(Type $var = null): Type
    {
    }

    /**
     * 保护方法，只有自身和子类可以访问。
     * @param Type
     * @return Type
     */
    protected function protectedFunction(Type $var = null): Type
    {
    }
  
    /**
     * 静态方法，可以在不实例化类的情况下执行。
     * @param Type
     * @return Type
     */
    public static function staticFunction(Type $var = null): Type
    {
    }

    // --> 魔术方法 <--

    /**
     * 具有构造函数的类会在每次创建新对象时先调用此方法，所以非常适合在使用对象之前做一些初始化工作。
     * http://php.net/manual/zh/language.oop5.decon.php
     * @param Type
     * @return void
     */
    public function __construct(Type $var = null)
    {
    }

    /**
     * 析构函数会在到某个对象的所有引用都被删除或者当对象被显式销毁时执行。
     * http://php.net/manual/zh/language.oop5.decon.php
     * @return void
     */
    public function __destruct()
    {
    }

    /**
     * 在给不可访问属性赋值时，__set() 会被调用。
     * http://php.net/manual/zh/language.oop5.overloading.php
     * @param string name
     * @param mixed value
     * @return void
     */
    public function __set(string $name , mixed $value)
    {
    }

    /**
     * 读取不可访问属性的值时，__get() 会被调用。
     * http://php.net/manual/zh/language.oop5.overloading.php
     * @param string name
     * @return mixed
     */
    public function __get(string $name)
    {
    }

    /**
     * 当对不可访问属性调用 isset() 或 empty() 时，__isset() 会被调用。
     * http://php.net/manual/zh/language.oop5.overloading.php
     * @param string name
     * @return bool
     */
    public function __isset(string $name)
    {
    }

    /**
     * 当对不可访问属性调用 unset() 时，__unset() 会被调用。
     * http://php.net/manual/zh/language.oop5.overloading.php
     * @param string name
     * @return void
     */
    public function __unset(string $name)
    {
    }

    /**
     * 在对象中调用一个不可访问方法时，__call() 会被调用。
     * http://php.net/manual/zh/language.oop5.overloading.php
     * @param string name
     * @param array arguments
     * @return mixed
     */
    public function __call(string $name, array $arguments)
    {
    }

    /**
     * 在静态上下文中调用一个不可访问方法时，__callStatic() 会被调用。
     * http://php.net/manual/zh/language.oop5.overloading.php
     * @param string name
     * @param array arguments
     * @return mixed
     */
    public static function __callStatic(string $name, array $arguments)
    {
    }

    /**
     * serialize() 函数会检查类中是否存在一个魔术方法 __sleep()。
     * 如果存在，该方法会先被调用，然后才执行序列化操作。此功能可以用于清理对象，
     * 并返回一个包含对象中所有应被序列化的变量名称的数组。
     * 如果该方法未返回任何内容，则 NULL 被序列化，并产生一个 E_NOTICE 级别的错误。
     * http://php.net/manual/zh/language.oop5.magic.php#object.sleep
     * @return array
     */
    public function __sleep()
    {
    }

    /**
     * 与之相反，unserialize() 会检查是否存在一个 __wakeup() 方法。
     * 如果存在，则会先调用 __wakeup 方法，预先准备对象需要的资源。
     * http://php.net/manual/zh/language.oop5.magic.php#object.wakeup
     * @return void
     */
    public function __wakeup()
    {
    }

    /**
     * __toString() 方法用于一个类被当成字符串时应怎样回应。
     * 例如 echo $obj; 应该显示些什么。此方法必须返回一个字符串，
     * 否则将发出一条 E_RECOVERABLE_ERROR 级别的致命错误。
     * http://php.net/manual/zh/language.oop5.magic.php#object.tostring
     * @return string
     */
    public function __toString()
    {
    }

    /**
     * 当尝试以调用函数的方式调用一个对象时，__invoke() 方法会被自动调用。
     * http://php.net/manual/zh/language.oop5.magic.php#object.invoke
     * @param Type
     * @return mixed
     */
    public function __invoke(Type $var = null)
    {
    }

    /**
     * 自 PHP 5.1.0 起当调用 var_export() 导出类时，此静态 方法会被调用。
     * http://php.net/manual/zh/language.oop5.magic.php#object.set-state
     * @param array properties
     * @return object
     */
    public static function __set_state(array $properties)
    {
    }

    /**
     * 在使用 var_dump() 时，会被调用。
     * http://php.net/manual/zh/language.oop5.magic.php#object.debuginfo
     * @return array
     */
    public function __debugInfo()
    {
    }
}

/**
 * 接口
 * 任何实现接口的类，都必须实现接口中的方法。
 */
interface InterfaceName
{

    public function FunctionName(Type $var = null): Type;

}

/**
 * 抽象类
 * 抽象类中可以包含普通方法，和抽象方法。
 */
abstract class AbstractClassName
{

    /**
     * 继承本抽象类的类，必须实现抽象方法。
     * @param Type
     * @return Type
     */
    abstract function abstractFunction(Type $var = null): Type;

}

/**
 * Trait
 * 提供代码复用能力、多继承能力
 */
trait Logger
{
	public function log($message)
	{
		return $message;
	}
}

class WriteLog
{
    use Logger;
  
    public function main()
    {
        return $this->log();
    }
}


/**
 * Enums (PHP >=8.1)
 * https://www.php.net/manual/zh/language.types.enumerations.php
 */

 interface StateCode {
    public function stateCode() : int;
 }

 enum States implements StateCode {
     case Running;
     case Stopped;

     public function stateCode() : int {
         return match($this) {
             State::Running => '444',
             State::Stopped => '666'
         };
     }
 }

 /**
  * You can also declare backed Enums
  */
  enum States : int implements StateCode {
    case Running = 1;
    case Stopped = 0;

    public function stateCode() : int {
        return match($this) {
            State::Running => '444',
            State::Stopped => '666'
        };
    }
}

 /** Enums can be use as a type */
 function notify(State $state) {
     // ...
 }
 notify(State::Running);

/**
 * PHP Regex.
 */

// Subpattern Modifiers & Assertions.
(?:)    Non capturing subpattern    ((?:foo|fu)bar) matches foobar or fubar without foo or fu appearing as a captured subpattern
(?=)    Positive look ahead assertion   foo(?=bar) matches foo when followed by bar
(?!)    Negative look ahead assertion   foo(?!bar) matches foo when not followed by bar
(?<=)   Positive look behind assertion  (?<=foo)bar matches bar when preceded by foo
(?<!)   Negative look behind assertion  (?<!foo)bar matches bar when not preceded by foo
(?>)    Once-only subpatterns   (?>\d+)bar Performance enhancing when bar not present
(?(x))  Conditional subpatterns (?(3)foo|fu)bar Matches foo if 3rd subpattern has matched, fu if not
(?#)    Comment (?# Pattern does x y or z)

// Base Character Classes
\w  匹配所有字母数字，等同于 [a-zA-Z0-9_]
\W  匹配所有非字母数字，即符号，等同于： [^\w]
\s 匹配所有空格字符，等同于： [\t\n\f\r\p{Z}]
\S  匹配所有非空格字符： [^\s]
\d  匹配数字： [0-9]
\D  匹配非数字： [^\d]
.   除换行符外的所有字符

// Multiplicity.
n*  	匹配>=0个重复的在*号之前的字符。
n+  	匹配>=1个重复的+号前的字符。
n?  	0或1个
{n} 	匹配n个
{n,}    只是匹配n个
{,m}    至多m个
{n,m}   匹配num个大括号之前的字符或字符集 (n <= num <= m).
(xyz)	字符集，匹配与 xyz 完全相等的字符串.
|			或运算符，匹配符号前或后的字符.
\			转义字符,用于匹配一些保留的字符 [ ] ( ) { } . * + ? ^ $ \ |
^			从开始行开始匹配.
$			从末端开始匹配.

// PHP Regular Expression Functions.

Function    Description
preg_match()    The preg_match() function searches string for pattern, returning true if pattern exists, and false otherwise.
preg_match_all()    The preg_match_all() function matches all occurrences of pattern in string. Useful for search and replace.
preg_replace()  The preg_replace() function operates just like ereg_replace(), except that regular expressions can be used in the pattern and replacement input parameters.
preg_split()    Preg Split (preg_split()) operates exactly like the split() function, except that regular expressions are accepted as input parameters.
preg_grep() The preg_grep() function searches all elements of input_array, returning all elements matching the regex pattern within a string.
preg_ quote()   Quote regular expression characters

// Code Snippets.

//filter_var.
//过滤准确邮箱
if (filter_var('test+email@fexample.com', FILTER_VALIDATE_EMAIL)) {
    echo "Your email is ok.";
} else {
    echo "Wrong email address format.";
}

//验证用户名，由字母数字（a-z，A-Z，0-9）和下划线组成，最少5个字符，最多20个字符。
$username = "user_name12";
if (preg_match('/^[a-z\d_]{5,20}$/i', $username)) {
    echo "Your username is ok.";
} else {
    echo "Wrong username format.";
}

//验证域名
$url = "http://domain-name.com/";
if (preg_match('/^(http|https|ftp):\/\/([A-Z0-9][A-Z0-9_-]*(?:\.[A-Z0-9][A-Z0-9_-]*)+):?(\d+)?\/?/i', $url)) {
    echo "Your url is ok.";
} else {
    echo "Wrong url.";
}

//从url提取域名 @与/作用一致
$url = "http://domain-name.com/index.html";
preg_match('/^(?:http://)?([^/]+)/i', $url, $matches);
$host = $matches[1];
echo $host; // domain-name.com

//Highlight a word in the content
//\b 是正则表达式规定的一个特殊代码（好吧，某些人叫它元字符，metacharacter），代表着单词的开头或结尾，也就是单词的分界处。
$text = "A regular expression (shortened as regex) is a sequence of characters that define a search pattern. Usually such patterns are used by string-searching algorithms for 'find' or 'find and replace' operations on strings, or for input validation.";
$text = preg_replace("/\b(regex)\b/i", 'replaced content', $text);
echo $text; /*A regular expression (shortened as replaced content) is a sequence of characters that define a search pattern. Usually such patterns are used by string-searching algorithms for 'find' or 'find and replace' operations on strings, or for input validation.*/
```
