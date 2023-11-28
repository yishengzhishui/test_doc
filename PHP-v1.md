## 1.google search

使用特定内容，调用google查询并返回结果
涉及的依赖：`GuzzleHttp`，`HtmlDomParser`

```php
<?php

declare(strict_types=1);
/**
 * This file is part of MangaToon server projects.
 */
namespace App\Library\CopyrightDetect\Webs\GoogleSearch;

use App\Library\CopyrightDetect\Webs\ProxyIp;
use GuzzleHttp\Client;
use GuzzleHttp\HandlerStack;
use Hyperf\Guzzle\CoroutineHandler;
use Mangatoon\Hyperf\Extend\Constants\Language;
use Mangatoon\Hyperf\Extend\Utils\LogUtil;
use Sunra\PhpSimple\HtmlDomParser;

class GoogleSearch
{
    private $keyword;

    private $search = 'https://www.google.com/search?q=';

    /**
     * @var string
     */
    private $url;

    private $language_code;

    public function __construct($keywords, $language = '')
    {
        $this->keyword = self::formatKeyword($keywords);
        $this->language_code = $language_code = is_numeric($language) ? Language::CODES[$language] : $language;

        $this->url = $this->search . $this->keyword;
        if ($language_code) {
            $this->url .= '&lr=lang_' . $language_code;
        }
    }

    public function search($proxy_ip = null): array
    {
    //google会对短时间内请求多次的ip进行限制，所有需要动态ip
        $options['proxy'] = $proxy_ip;
        $content = $this->request('GET', $this->url, $options)->getBody()->getContents();
        if (empty($content)) {
            throw new \RuntimeException('content is empty');
        }
        return $this->formatResult($content);
    }

    public function request($method, $url, $options)
    {
        $handler_stack = HandlerStack::create(new CoroutineHandler());
        $config = ['handler' => $handler_stack];
        $client = new Client($config);
        return $client->request($method, $url, $options);
    }

    public function formatResult($html)
    {
        $html = HtmlDomParser::str_get_html($html);
        file_put_contents(BASE_PATH . '/runtime/temp_test.html', $html);

        $nodes = $html->find('.Gx5Zad');
        if (empty($nodes)) {
            LogUtil::stdout()->debug('content is valid html file');
            return [];
        }
        $nodes = $html->find('.Gx5Zad');

        $return = [];
        foreach ($nodes as $node) {
            try {
                $f = $node->find('h3');
                if (empty($f)) {
                    continue;
                }
                $item = new \stdClass();
                $item->title = $node->find('h3[class=zBAuLc]', 0)->find('.BNeawe', 0)->innertext();
                $item->url = html_entity_decode(str_replace('/url?q=', '', $node->find('a', 0)->href));
                $return[] = $item;
            } catch (\RuntimeException $e) {
                echo $e->getMessage() . PHP_EOL;
            }
        }
        return $return;
    }

    public static function formatKeyword($keywords)
    {
        return implode('+', $keywords);
    }
}
```

## 2.页面加载并获取页面内容

`Browsershot` 模拟谷歌浏览器进行加载

```shell
// mac
npm install puppeteer
composer require spatie/browsershot

//centos
npm install puppeteer --location=global
sudo yum install -y chromium
sudo yum -y install libXScrnSaver-1.2.2-6.1.el7.x86_64
```

```php
    public function getBuenovelaContentChapterAction($url = 'https://www.buenovela.com/book/31000385032-3201435')
    {
        try {
            if (ENV::isDev()) {
            // 本地环境，需要vpn转发需要制定代理
            // waitUntilNetworkIdle 等待所有页面渲染加载完成
                $content = Browsershot::url($url)->setProxyServer('127.0.0.1:1087')->waitUntilNetworkIdle()->bodyHtml();
            } else {
            // 在Linux机器上 需要安装chromium-browser依赖，并且指定参数
                $content = Browsershot::url($url)->setChromePath('/usr/bin/chromium-browser')->addChromiumArguments(['no-sandbox', 'disable-setuid-sandbox'])->waitUntilNetworkIdle()->bodyHtml();
            }
            $html = HtmlDomParser::str_get_html($content);
        } catch (\RuntimeException $e) {
            $this->info('页面异常 跳过-1');
            return ['url' => $url];
        }

        if (!empty($html->find('.chapter_content', 0))) {
            $read_chapter = $html->find('.scroll_content', 0)->find('h1', 0)->innertext();
            $contents = $html->find('.chapter_content', 0)->find('p');
            $context_arr = [];
            foreach ($contents as $content) {
                $p = $content->innertext();
                if (!empty($p)) {
                    $context_arr[] = $p;
                }
            }
            if (empty($context_arr)) {
                $context_arr = explode('.', $html->find('.chapter_content', 0)->text());
            }
            return ['read_chapter' => $read_chapter, 'context' => $context_arr, 'url' => $url];
        }

        $nodes = $html->childNodes();
        $context_arr = [];
        foreach ($nodes as $node) {
            $content = explode('.', $node->text());
            $context_arr = array_merge($context_arr, $content);
        }

        return ['read_chapter' => '', 'context' => $context_arr, 'url' => $url];
    }
```

## 3.文件夹下所有文件

```php
private function getAllFiles($path)
    {
        $files = [];
        $handler = opendir($path);
        while (($filename = readdir($handler)) !== false) {
            if ($filename !== '.' && $filename !== '..' && $filename !== '.DS_Store' && strpos($filename, '.') !== false) {
                $files[] = $filename;
            }
        }
        closedir($handler);
        return $files;
    }
```

## 4.File 文件处理相关

```php
<?php

declare(strict_types=1);
/**
 * This file is part of MangaToon server projects.
 */
namespace Mangatoon\Hyperf\Extend\Utils;

use DirectoryIterator;
use Hyperf\Utils\Str;
use RuntimeException;

class File
{
    public const IMAGE_SUFFIXES = [
        'jpg', 'jpeg', 'png', 'gif', 'bmp', 'ico', 'tif',
    ];

    public const COMPRESSED_ARCHIVE_SUFFIXES = [
        'zip', 'tar.gz', 'tar.bz2', 'gz', 'rar', 'tar', 'tsv.gz', 'csv.gz', 'log.gz',
    ];

    /**
     * 从文件名中获取一个文件的后缀
     * @param string $filename
     * @return string
     */
    public static function getSuffixOfFilename($filename)
    {
        if (preg_match('/(\\.[a-z0-9A-Z]+){1,2}$/', $filename, $matches)) {
            $suffix = substr($matches[0], 1);
            // 以下代码是为了兼容用户使用 1-1.2.jpg 这种命名方式的情况
            if (in_array($suffix, self::COMPRESSED_ARCHIVE_SUFFIXES)) {
                return $suffix;
            }
            if (preg_match('/\\.([a-z0-9A-Z]+)$/', $suffix, $matches)) {
                return $matches[1];
            }
            return $suffix;
        }

        return '';
    }

    /**
     * 判断一个文件后缀是不是图片.
     * @param string $suffix
     * @return bool
     */
    public static function isFileSuffixImage($suffix)
    {
        return in_array(strtolower($suffix), self::IMAGE_SUFFIXES);
    }

    public static function sortFileNames($file_names)
    {
        $file_names_sorted = [];
        foreach ($file_names as $file_name) {
            $file_name_for_sort = str_replace(
                ['一', '二', '三', '四', '五', '六', '七', '八', '九', '十'],
                ['1', '2', '3', '4', '5', '6', '7', '8', '9', '10'],
                $file_name
            );
            $num = preg_match('/([0-9]+)/', $file_name_for_sort, $matches) ? intval($matches[1]) : PHP_INT_MAX;
            $file_names_sorted[$file_name] = ['num' => $num, 'name' => $file_name];
        }
        uasort($file_names_sorted, function ($a, $b) {
            if ($a['num'] === PHP_INT_MAX && $b['num'] === PHP_INT_MAX) {
                return strcasecmp($a['name'], $b['name']);
            }
            if ($a['num'] === $b['num']) {
                return strcasecmp($a['name'], $b['name']);
            }
            return $a['num'] - $b['num'];
        });

        return array_keys($file_names_sorted);
    }

    /**
     * 计算一个文件或者目录的字节大小.
     * @param string $file_path 文件或者目录的路径
     * @return int
     */
    public static function calculateFileOrFolderSize($file_path)
    {
        $file_size = 0;
        if (is_file($file_path)) {
            $file_size += filesize($file_path);
        } elseif (is_dir($file_path)) {
            if ($file_path[-1] !== '/') {
                $file_path .= '/';
            }
            $dir = opendir($file_path);
            while ($d = readdir($dir)) {
                if ($d == '.' || $d == '..') {
                    continue;
                }
                $sub_file_path = $file_path . $d;
                $file_size += self::calculateFileOrFolderSize($sub_file_path);
            }
        }

        return $file_size;
    }

    /**
     * @param string $path the folder path
     * @param mixed $extension string or array, if specified, only files with theses extensions will be listed
     * @param int $depth the depth, default -1 means not limit depth of search, 1 means only current and no sub folders
     * @return \SplFileInfo[]
     */
    public static function listAllFilesWithExtensionUnder($path, $extension = null, $depth = -1)
    {
        if (!is_dir($path)) {
            return [];
        }
        if (is_string($extension)) {
            $extensions = [strtolower($extension)];
        } elseif (is_array($extension)) {
            $extensions = [];
            foreach ($extension as $item) {
                $extensions[] = strtolower($item);
            }
        } else {
            $extensions = [];
        }

        $files = [];
        $it = new DirectoryIterator($path);
        while ($it->valid()) {
            if ($it->isDot()) {
            } elseif ($it->isFile()) {
                if (empty($extensions)) {
                    $files[] = $it->getFileInfo();
                } elseif (in_array(strtolower($it->getExtension()), $extensions)) {
                    $files[] = $it->getFileInfo();
                }
            } elseif ($it->isDir()) {
                if ($depth != 0) {
                    $subpath = $it->getPath() . DIRECTORY_SEPARATOR . $it->getFilename();
                    $subfiles = self::listAllFilesWithExtensionUnder($subpath, $extensions, $depth - 1);
                    foreach ($subfiles as $item) {
                        $files[] = $item;
                    }
                }
            }
            $it->next();
        }

        return $files;
    }

    /**
     * @param int $size
     * @return string
     */
    public static function formatFileSize($size)
    {
        if ($size < 1024) {
            return '<1K';
        }
        if ($size < 1024 * 1024) {
            return (round($size * 10 / 1024) / 10) . 'K';
        }
        if ($size < 1024 * 1024 * 1024) {
            return (round($size * 10 / 1024 / 1024) / 10) . 'M';
        }
        return (round($size * 10 / 1024 / 1024 / 1024) / 10) . 'G';
    }

    /**
     * 判断一个目录是否是空目录.
     * @param $dir
     * @return bool
     */
    public static function isDirEmpty($dir)
    {
        if (!is_readable($dir)) {
            throw new RuntimeException('cannot read dir ' . $dir);
        }
        if (!is_dir($dir)) {
            throw new RuntimeException($dir . ' is not dir');
        }
        return count(scandir($dir)) == 2;
    }

    /**
     * 从一批文件路径中找到一个存在的可执行文件.
     * @param string[] $folder_paths
     * @return string
     */
    public static function findExecutableInFolderPaths(string $name, array $folder_paths): ?string
    {
        foreach ($folder_paths as $folder_path) {
            if (empty($folder_path)) {
                continue;
            }
            $file_path = $folder_path[strlen($folder_path) - 1] === '/'
                ? $folder_path . $name
                : $folder_path . '/' . $name;
            if (is_executable($file_path)) {
                return $file_path;
            }
        }
        return null;
    }

    /**
     * 在默认的系统目录下找到一个可执行文件.
     * @return string
     */
    public static function findExecutable(string $name): ?string
    {
        $folder_paths = array_merge(explode(':', $_ENV['PATH'] ?? ''), [
            '/Applications/MAMP/Library/bin/',
            '/usr/local/bin/',
            '/usr/bin/',
            '/usr/sbin/',
        ]);
        return self::findExecutableInFolderPaths($name, $folder_paths);
    }

    /**
     * 下载一个 URL 地址到本地某个文件.
     */
    public static function download(string $url, string $file_path): bool
    {
        $axel = self::findExecutable('axel');
        $command = "rm -f '{$file_path}'; {$axel} -n 10 -o '{$file_path}' '{$url}'";
        if (empty($axel)) {
            if (PHP_OS === 'Darwin') {
                LogUtil::stdout()->error('Please install axel first, you can use command: brew install axel');
                throw new \RuntimeException('axel command not found');
            }
            $command = "rm -f '{$file_path}'; wget -O '{$file_path}' '{$url}'";
        }
        @mkdir(dirname($file_path), 0777, true);
        ExecUtil::execAndCheckReturnVar($command);
        return true;
    }

    /**
     * 读取一个文件，并对每一行内容执行回调函数.
     */
    public static function foreachLine(string $file_path, callable $callback)
    {
        $is_gz = Str::endsWith($file_path, '.gz');
        $handle = $is_gz ? gzopen($file_path, 'r') : fopen($file_path, 'r');
        while (true) {
            $line = $is_gz ? gzgets($handle) : fgets($handle);
            if ($line === false) {
                break;
            }
            if (empty($line)) {
                continue;
            }
            $callback($line);
        }
        $is_gz ? gzclose($handle) : fclose($handle);
    }

    /**
     * 把图片压制为 webp 格式.
     * @param string $from
     * @param string $to
     * @return bool
     */
    public static function convertImageToWebp($from, $to)
    {
        $cwebp = '/usr/local/bin/cwebp';
        if (file_exists('/usr/bin/cwebp')) {
            $cwebp = '/usr/bin/cwebp';
        }
        if (self::isFileFormatGif($from)) {
            $cwebp = '/usr/local/bin/gif2webp';
            if (file_exists('/usr/bin/gif2webp')) {
                $cwebp = '/usr/bin/gif2webp';
            }
        }
        $command = "'{$cwebp}' '{$from}' -o '{$to}'";
        exec($command, $output, $result);

        return $result === 0 && is_file($to);
    }

    /**
     * 判断一个文件名是不是 gif 格式.
     * @param $file_name
     * @return bool
     */
    public static function isFileFormatGif($file_name)
    {
        return !empty(@imagecreatefromgif($file_name));
    }

    /**
     * 判断一个文件名是不是 png 格式.
     * @param $file_name
     * @return bool
     */
    public static function isFileFormatPng($file_name)
    {
        return !empty(@imagecreatefrompng($file_name));
    }

    /**
     * 优化 png 文件.
     * @param $png_file_path
     * @return bool
     */
    public static function optimizeImagePNG($png_file_path)
    {
        $pngquant = '';
        if (is_executable('/usr/bin/pngquant')) {
            $pngquant = '/usr/bin/pngquant';
        }
        if (empty($pngquant) && is_executable('/usr/local/pngquant/pngquant')) {
            $pngquant = '/usr/local/pngquant/pngquant';
        }
        if (is_executable($pngquant)) {
            $pngquant_temp_path = str_replace('.png', '', $png_file_path) . '-' . time() . '-pngquant.png';
            $command = "{$pngquant} --quality=65-80 --strip --output '{$pngquant_temp_path}' -- '{$png_file_path}'";
            exec($command, $output, $result);
            if ($result === 0) {
                @copy($pngquant_temp_path, $png_file_path);
                @unlink($pngquant_temp_path);
                return true;
            }
        }

        //可以进一步压缩
        $result = null;
        $optipng = '/usr/bin/optipng';
        if (is_executable('/usr/local/bin/optipng')) {
            $optipng = '/usr/local/bin/optipng';
        }
        $command = "{$optipng} -strip all -quiet -o3 '{$png_file_path}'";
        exec($command, $output, $result);
        if ($result === 0) {
            return true;
        }
        return false;
    }

    /**
     * 优化 jpg 文件.
     * @param string $from
     * @param string $to
     * @return bool
     */
    public static function optimizeImageJPG($from, $to)
    {
        $jpegtran = '/usr/local/bin/jpegtran';
        if (file_exists('/usr/bin/jpegtran')) {
            $jpegtran = '/usr/bin/jpegtran';
        }
        $command = "{$jpegtran} -copy none -optimize -progressive -outfile '{$to}' '{$from}'";
        exec($command, $output, $result);

        return $result === 0 && is_file($to);
    }

    /**
     * 判断一个文件名是不是 jpg 格式.
     * @param $file_name
     * @return bool
     */
    public static function isFileFormatJpg($file_name)
    {
        return !empty(@imagecreatefromjpeg($file_name));
    }
}

```


## 5.Http 网络处理相关


```php
<?php

declare(strict_types=1);
/**
 * This file is part of MangaToon server projects.
 */
namespace Mangatoon\Hyperf\Extend\Utils;

use GuzzleHttp\Client;
use GuzzleHttp\HandlerStack;
use GuzzleHttp\RequestOptions;
use Hyperf\Guzzle\CoroutineHandler;
use Hyperf\Utils\ApplicationContext;
use Mangatoon\Hyperf\Extend\Compatibility\Compatibility;
use Mangatoon\Hyperf\Extend\Compatibility\YafProjectPerformanceTracer;
use Mangatoon\Hyperf\Extend\Events\HttpRequestCompleted;
use Psr\EventDispatcher\EventDispatcherInterface;
use Psr\Http\Message\ResponseInterface;
use Swoole\Coroutine;

class Http
{
    public const DEFAULT_HEADERS = [
        'Accept' => 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
        'Accept-Charset' => 'GB2312,utf-8;q=0.7,*;q=0.7',
        'Accept-Encoding' => 'gzip,deflate',
        'User-Agent' => 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3381.0 Safari/537.36',
    ];

    public static function get($url, $headers = [], $follow = true, $timeout = 5, $sink = null): ResponseInterface
    {
        $client = self::createGuzzleClient();
        $options = [
            RequestOptions::HEADERS => array_replace(self::DEFAULT_HEADERS, $headers),
            RequestOptions::TIMEOUT => $timeout,
            RequestOptions::ALLOW_REDIRECTS => $follow,
        ];
        if ($sink) {
            $options[RequestOptions::SINK] = $sink;
        }
        [$url, $options] = self::resolveAuthInUrl($url, $options);
        $options = self::addProxyToOptions($url, $options);
        $response = $client->get($url, $options);
        return $response;
    }

    public static function post($url, $data, $headers = [], $follow = true, $timeout = 5, $sink = null): ResponseInterface
    {
        $body = (is_array($data)) ? http_build_query($data) : $data;
        $client = self::createGuzzleClient();
        $options = [
            RequestOptions::HEADERS => array_replace(self::DEFAULT_HEADERS, [
                'Content-Type' => 'application/x-www-form-urlencoded',
            ], $headers),
            RequestOptions::TIMEOUT => $timeout,
            RequestOptions::ALLOW_REDIRECTS => $follow,
            RequestOptions::BODY => $body,
        ];
        if ($sink) {
            $options[RequestOptions::SINK] = $sink;
        }
        [$url, $options] = self::resolveAuthInUrl($url, $options);
        $options = self::addProxyToOptions($url, $options);
        $response = $client->post($url, $options);
        return $response;
    }

    /**
     * @param string $method HTTP Method, 支持 delete\put
     * @param $url
     * @param $data
     * @param array $headers
     * @param bool $follow
     * @param int $timeout
     * @param null $sink
     * @return ResponseInterface
     */
    public static function httpRequestByMethod(string $method, $url, $data, $headers = [], $follow = true, $timeout = 5, $sink = null): ResponseInterface
    {
        $client = self::createGuzzleClient();
        $options = [
            RequestOptions::HEADERS => array_replace(self::DEFAULT_HEADERS, [
                'Content-Type' => 'application/x-www-form-urlencoded',
            ], $headers),
            RequestOptions::TIMEOUT => $timeout,
            RequestOptions::ALLOW_REDIRECTS => $follow,
            RequestOptions::BODY => (is_array($data)) ? http_build_query($data) : $data,
        ];
        if ($sink) {
            $options[RequestOptions::SINK] = $sink;
        }
        [$url, $options] = self::resolveAuthInUrl($url, $options);
        $options = self::addProxyToOptions($url, $options);
        return $client->$method($url, $options);
    }

    public static function createGuzzleClient($options = []): Client
    {
        if (Compatibility::isHyperf() && Coroutine::getCid() > 0) {
            $handler_stack = HandlerStack::create(new CoroutineHandler());
            $handler_stack->push(self::getEventHandler());
            $options = array_merge(['handler' => $handler_stack], $options);
            return make(Client::class, $options);
        }
        return new Client($options);
    }

    private static function getEventHandler(): callable
    {
        return function (callable $handler) {
            return function ($request, array $options) use ($handler) {
                $request_start_time = microtime(true);
                return $handler($request, $options)
                    ->then(function ($response) use ($request, $request_start_time) {
                        /** @var \GuzzleHttp\Psr7\Request $request */
                        /** @var \GuzzleHttp\Psr7\Response $response */
                        $event = new HttpRequestCompleted($request, $response);
                        $event->time = (microtime(true) - $request_start_time) * 1000;
                        /** @var EventDispatcherInterface $dispatcher */
                        $dispatcher = ApplicationContext::getContainer()->get(EventDispatcherInterface::class);
                        $dispatcher->dispatch($event);
                        return $response;
                    });
            };
        };
    }

    private static function resolveAuthInUrl($url, $options)
    {
        // 对 URL 地址中包含了用户登录验证的情况做兼容
        $uri = parse_url($url);
        if (!empty($uri['user'])) {
            $options[RequestOptions::AUTH] = [$uri['user'], $uri['pass']];
            $url = $uri['scheme'] . '://' . $uri['host'];
            if (!empty($uri['port'])) {
                $url .= ':' . $uri['port'];
            }
            if (!empty($uri['query'])) {
                $url .= '?' . $uri['query'];
            }
        }
        return [$url, $options];
    }

    private static function checkLocalProxy()
    {
        // ShadowSocks : 1080 | 1087, Clash : 7890
        foreach ([1080, 1087, 7890] as $port) {
            $sock = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
            $res = socket_connect($sock, 'localhost', $port);
            socket_close($sock);
            if ($res) {
                return $port;
            }
        }
        return false;
    }

    private static function addProxyToOptions($url, $options)
    {
        if (in_array(parse_url($url)['host'], [
            'graph.facebook.com',
            'accounts.google.com',
            'admob.googleapis.com',
            'www.googleapis.com',
            'oauth2.googleapis.com',
            'fcm.googleapis.com',
            'storage.googleapis.com',
        ])) {
            if (ENVUtil::isDev() && $port = self::checkLocalProxy()) {
                $options[RequestOptions::PROXY] = 'http://localhost:' . $port;
            } elseif (ENVUtil::isTest()) {
                $options['swoole'] = [
                    'socks5_host' => '127.0.0.1',
                    'socks5_port' => 1086,
                ];
            }
        }
        return $options;
    }
}

```
