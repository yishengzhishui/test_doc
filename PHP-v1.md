1.google search
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

2.页面加载并获取页面内容
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

