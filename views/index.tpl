<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport"
        content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>朋也接口</title>

  <link href="//cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
  <style>
    .panel {
      border-top: 0;
    }
  </style>

  <script src="//cdn.bootcss.com/jquery/2.2.2/jquery.min.js"></script>
  <script src="//cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</head>
<body style="padding-top: 20px;">
<div class="container">
  <div class="jumbotron">
    <h1>朋也接口</h1>
    <p>解析网页整理的接口，用于开发app，严禁用于其它用途</p>
  </div>

  <ul class="nav nav-tabs" role="tablist">
    <li role="presentation" class="active">
      <a href="#poxiao" aria-controls="poxiao" role="tab" data-toggle="tab">破晓电影</a>
    </li>
    <li role="presentation">
      <a href="#piaohua" aria-controls="piaohua" role="tab" data-toggle="tab">飘花电影</a>
    </li>
    <li role="presentation">
      <a href="#hldm" aria-controls="hldm" role="tab" data-toggle="tab">红旅动漫</a>
    </li>
  </ul>

  <div class="tab-content">
    <div class="panel panel-default tab-pane active" role="tabpanel" id="poxiao">
      <div class="panel-body">
        <h3>首页</h3>
        <table class="table table-bordered">
          <tr>
            <td>地址</td>
            <td>/poxiao/index</td>
          </tr>
          <tr>
            <td>参数</td>
            <td>无</td>
          </tr>
          <tr>
            <td>例子</td>
            <td><a href="/poxiao/index" target="_blank">https://apis.tomoya.cn/poxiao/index</a></td>
          </tr>
          <tr>
            <td>对应原网页</td>
            <td><a href="http://www.poxiao.com" target="_blank">http://www.poxiao.com</a></td>
          </tr>
        </table>
        <hr>

        <h3>电影详情</h3>
        <table class="table table-bordered">
          <tr>
            <td>地址</td>
            <td>/poxiao/detail</td>
          </tr>
          <tr>
            <td>参数</td>
            <td>url: 电影详情页面链接</td>
          </tr>
          <tr>
            <td>例子</td>
            <td><a href="/poxiao/detail?url=http://www.poxiao.com/movie/42245.html" target="_blank">https://apis.tomoya.cn/poxiao/index?url=http://www.poxiao.com/movie/42245.html</a></td>
          </tr>
          <tr>
            <td>对应原网页</td>
            <td><a href="http://www.poxiao.com/movie/42245.html" target="_blank">http://www.poxiao.com/movie/42245.html</a></td>
          </tr>
        </table>
        <hr>

        <h3>电影大全</h3>
        <table class="table table-bordered">
          <tr>
            <td>地址</td>
            <td>/poxiao/movie</td>
          </tr>
          <tr>
            <td>参数</td>
            <td>pageNum: 页数</td>
          </tr>
          <tr>
            <td>例子</td>
            <td><a href="/poxiao/movie/?pageNum=1" target="_blank">https://apis.tomoya.cn/poxiao/movie/?pageNum=1</a></td>
          </tr>
          <tr>
            <td>对应原网页</td>
            <td><a href="http://www.poxiao.com/type/movie/" target="_blank">http://www.poxiao.com/type/movie/</a></td>
          </tr>
        </table>
      </div>
    </div>
    <div class="panel panel-default tab-pane" role="tabpanel" id="piaohua">
      <div class="panel-body">
        <h3>首页</h3>
        <table class="table table-bordered">
          <tr>
            <td>地址</td>
            <td>/poxiao/index</td>
          </tr>
          <tr>
            <td>参数</td>
            <td>无</td>
          </tr>
          <tr>
            <td>例子</td>
            <td><a href="/piaohua/index/" target="_blank">https://apis.tomoya.cn/piaohua/index/</a></td>
          </tr>
          <tr>
            <td>对应原网页</td>
            <td><a href="http://www.piaohua.com/" target="_blank">http://www.piaohua.com/</a></td>
          </tr>
        </table>
        <hr>

        <h3>电影详情</h3>
        <table class="table table-bordered">
          <tr>
            <td>地址</td>
            <td>/piaohua/detail</td>
          </tr>
          <tr>
            <td>参数</td>
            <td>url: 电影详情页面链接</td>
          </tr>
          <tr>
            <td>例子</td>
            <td><a href="/piaohua/detail?url=http://www.piaohua.com/html/xiju/2017/0621/32105.html" target="_blank">https://apis.tomoya.cn/piaohua/detail?url=http://www.piaohua.com/html/xiju/2017/0621/32105.html</a></td>
          </tr>
          <tr>
            <td>对应原网页</td>
            <td><a href="http://www.piaohua.com/html/xiju/2017/0621/32105.html" target="_blank">http://www.piaohua.com/html/xiju/2017/0621/32105.html</a></td>
          </tr>
        </table>
        <hr>

        <h3>电影大全</h3>
        <table class="table table-bordered">
          <tr>
            <td>地址</td>
            <td>/piaohua/movie</td>
          </tr>
          <tr>
            <td>参数</td>
            <td>pageNum: 页数, type: 电影类型（dongzuo, xiju, aiqing, kehuan, juqing, xuannian, wenyi, zhanzheng, kongbu, zainan, lianxuju, dongman, zongyijiemu, 3D, jilupian, huaijiu）</td>
          </tr>
          <tr>
            <td>例子</td>
            <td><a href="/piaohua/movie/?pageNum=1&type=dongzuo" target="_blank">https://apis.tomoya.cn/piaohua/movie/?pageNum=1&type=dongzuo</a></td>
          </tr>
          <tr>
            <td>对应原网页</td>
            <td><a href="http://www.piaohua.com/html/dongzuo/index.html" target="_blank">http://www.piaohua.com/html/dongzuo/index.html</a></td>
          </tr>
        </table>
        <hr>

        <h3>今日更新电影</h3>
        <table class="table table-bordered">
          <tr>
            <td>地址</td>
            <td>/piaohua/lastest</td>
          </tr>
          <tr>
            <td>参数</td>
            <td>无</td>
          </tr>
          <tr>
            <td>例子</td>
            <td><a href="/piaohua/lastest" target="_blank">https://apis.tomoya.cn/piaohua/lastest</a></td>
          </tr>
          <tr>
            <td>对应原网页</td>
            <td><a href="http://www.piaohua.com/html/zuixindianying.html" target="_blank">http://www.piaohua.com/html/zuixindianying.html</a></td>
          </tr>
        </table>
      </div>
    </div>
    <div class="panel panel-default tab-pane" role="tabpanel" id="hldm">
      <div class="panel-body">
        <span class="glyphicon glyphicon-warning-sign text-danger"></span> 红旅动漫网站内容好像2017年都没有更新了，暂时不解析了
      </div>
    </div>
  </div>
</div>
<hr>
<div class="footer text-center">
  <p><a href="https://github.com/tomoya92/pyapis">源码</a> &nbsp; <a href="https://github.com/tomoya92/pyapis/issues">有问题？</a></p>
  <p>
    本站总访问量<span id="busuanzi_value_site_pv"></span>次，本站访客数<span id="busuanzi_value_site_uv"></span>人次，本文总阅读量<span id="busuanzi_value_page_pv"></span>次
  </p>
  <script async src="https://dn-lbstatics.qbox.me/busuanzi/2.3/busuanzi.pure.mini.js"></script>
</div>
</body>
</html>