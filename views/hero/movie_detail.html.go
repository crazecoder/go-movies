// Code generated by hero.
// source: E:\www\project\dexter\open\go-movies\views\hero\movie_detail.html
// DO NOT EDIT!
package template

import (
	"bytes"
	"go_movies/services"
	"go_movies/utils"
	"time"

	"github.com/shiyanhui/hero"
)

func MDetail(show map[string]interface{}, buffer *bytes.Buffer) {
	buffer.WriteString(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head profile="http://gmpg.org/xfn/11">
    <meta charset="UTF-8">
    <meta http-equiv="Content-Type" content="text/html">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <title>GoMovies-电影</title>
    <meta name="keywords" content="GoMovies-电影">
    <meta name="description" content="GoMovies-电影">
    <meta property="wb:webmaster" content="bec25808">
    <meta name="referrer" content="never">
    <link rel="canonical" href="/static/css/css">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=0, minimum-scale=1.0, maximum-scale=1.0">
    <link rel="stylesheet" type="text/css" href="/static/css/kube.css">
    <link rel="stylesheet" type="text/css" href="/static/css/reset.css">
    <link rel="stylesheet" type="text/css" href="/static/css/style.css">
    <link rel="stylesheet" type="text/css" href="//cdn.bootcss.com/font-awesome/4.7.0/css/font-awesome.min.css">
    <link rel="shortcut icon" href="/static/image/favicon.ico" type="image/x-icon" />
    <script>var killIE6ImgUrl = "/skin/66scc/images";</script>
    <!--<link rel="stylesheet" type="text/css" href="/static/css/lightbox.css" />-->
    <!--<script type='text/javascript' src='/static/js/lightbox.min.js'></script>-->
    <!--[if lt IE 9]>
    <script src="/static/js/html5.js"></script>
    <![endif]-->
</head>
`)
	buffer.WriteString(`

<body class="custom-background">


`)
	buffer.WriteString(`
    <div id="head" class="row">
        <div class="container row">
            <div class="row">
                <div id="topbar">
                    <ul id="toolbar" class="menu">

                    </ul>
                </div>
                <div id="rss">
                    <ul>
                        <li><a href="javascript:;"  title="最新更新文字版">不要点击</a>  </li>
                    </ul>
                </div>
            </div>
        </div>
        <div class="clear"></div>
        <div class="container">
            <div id="blogname" class="third"> <a href="/" title="6v电影-新版">
                    <h1>
                        go-movies      </h1>
                    <img src="/static/image/logo_.png" alt="6v电影-新版"></a> </div>

        </div>
        <div class="clear"></div>
    </div>`)
	buffer.WriteString(`<div class="mainmenus container" id="nav_b">
        <div class="mainmenu">
            <div class="topnav">
                <ul id="menus">

                    <li class="menu-item
                    `)
	if show["nav_link"] == "/" {
		hero.EscapeHTML("current_page_item", buffer)
	}
	buffer.WriteString(` ">
                    <a href="/">首页</a>
                    </li>


                    `)
	for _, category := range show["categories"].([]utils.Categories) {
		buffer.WriteString(`
                        <li class="menu-item
                        `)
		if category.Link == show["nav_link"] {
			hero.EscapeHTML("current_page_item", buffer)
		}
		buffer.WriteString(` ">
                    <a href="/?cate=`)
		hero.EscapeHTML(category.Link, buffer)
		buffer.WriteString(` ">`)
		hero.EscapeHTML(category.Name, buffer)
		buffer.WriteString(`</a>

                    <ul class="sub-menu">
                        `)
		for _, subCategory := range category.Sub {
			buffer.WriteString(`
                        <li class="menu-item">
                            <a href="/?cate=`)
			hero.EscapeHTML(subCategory.Link, buffer)
			buffer.WriteString(`">`)
			hero.EscapeHTML(subCategory.Name, buffer)
			buffer.WriteString(`</a>
                        </li>
                        `)
		}
		buffer.WriteString(`
                    </ul>


                    </li>
                    `)
	}
	buffer.WriteString(`

                    <li class="menu-item
                    `)
	if show["nav_link"] == "/about" {
		hero.EscapeHTML("current_page_item", buffer)
	}
	buffer.WriteString(` ">
                    <a href="/about"> 关于 </a>
                    </li>

                </ul>
                <div id="select_menu">
                    <select onChange="document.location.href=this.options[this.selectedIndex].value;" id="select-menu-nav">
                        <option value="#">导航菜单</option>
                        `)
	for _, category := range show["categories"].([]utils.Categories) {
		buffer.WriteString(`
                        <option   value="/?cate=`)
		hero.EscapeHTML(category.Link, buffer)
		buffer.WriteString(` ">`)
		hero.EscapeHTML(category.Name, buffer)
		buffer.WriteString(`</option>
                        `)
		for _, subCategory := range category.Sub {
			buffer.WriteString(`
                        <option   value="/?cate=`)
			hero.EscapeHTML(subCategory.Link, buffer)
			buffer.WriteString(` ">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`)
			hero.EscapeHTML(subCategory.Name, buffer)
			buffer.WriteString(`</option>
                        `)
		}
	}
	buffer.WriteString(`
                        <option   value="about"> 关于 </option>
                    </select>
                </div>
            </div>
        </div>
        <div class="clear"></div>
    </div>


<div class="search_m box row">
    <div class="search_site" style="width: 100%;">
        <form method="get" name="searchform" id="searchform" action="/search">
            <input type="submit" value="" id="searchsubmit" class="button">
            <label><span>请输入搜索内容，仅支持电影名称</span>
                <input type="text" class="search-s" name="q" required>
            </label>
        </form>
    </div>
</div>
`)
	buffer.WriteString(`


	<div class="container">
        		        <div class="row">
			                <div class="subsidiary box">
                    <div class="bulletin fourfifth">
                        <span class="sixth">当前位置：</span>       `)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["detail"].(map[string]interface{})["type"].(string), buffer)
	buffer.WriteString(`              </div>
                </div>
                    </div>



           `)
	buffer.WriteString(`
    <div id="sidebar">


        <div class="search box row">
            <div class="search_site">
                <form method="get" name="searchform" id="searchform" action="/search">
                    <input type="submit" value="" id="searchsubmit" class="button">
                    <label><span>请输入搜索内容，仅支持电影名称</span>
                        <input type="text" class="search-s" name="q" required>
                    </label>
                </form>
            </div>
        </div>

        <div class="widget box row">
            <div id="tab-title">
                <div class="tab">
                    <ul id="tabnav">
                        <li class="selected">最新电影</li>
                        <li class="">最新剧集</li>
                        <li class="">站长推荐</li>
                    </ul>
                </div>
                <div class="clear"></div>
            </div>
            <div id="tab-content">
                <ul class="">
                    `)
	for _, new_film := range show["new_film"].([]services.MovieListStruct) {
		buffer.WriteString(`
                    <li>
                        <a href='/movie?link=`)
		hero.EscapeHTML(new_film.Link, buffer)
		buffer.WriteString(`'>
                            `)
		hero.EscapeHTML(new_film.Name, buffer)
		buffer.WriteString(`
                        </a>
                    </li>
                    `)
	}
	buffer.WriteString(`

                </ul>
                <ul class="hide">
                    `)
	for _, new_tv := range show["new_tv"].([]services.MovieListStruct) {
		buffer.WriteString(`
                    <li>
                        <a href='/movie?link=`)
		hero.EscapeHTML(new_tv.Link, buffer)
		buffer.WriteString(`'>
                            `)
		hero.EscapeHTML(new_tv.Name, buffer)
		buffer.WriteString(`
                        </a>
                    </li>
                    `)
	}
	buffer.WriteString(`
                </ul>
                <ul class="hide">
                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-26855.html'>
                            表姐，你好嘢！
                        </a>
                    </li>

                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-24088.html'>
                            维多利亚一号
                        </a>
                    </li>

                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-14876.html'>
                            爱人同志
                        </a>
                    </li>

                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-26856.html'>
                            表姐，你好嘢！2
                        </a>
                    </li>

                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-26857.html'>
                            表姐，你好嘢！3之大人驾到
                        </a>
                    </li>

                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-26858.html'>
                            表姐，你好嘢！4之情不自禁
                        </a>
                    </li>

                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-14334'>
                            绝命毒师第一季完结
                        </a>
                    </li>
                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-14335.html'>
                            绝命毒师第二季完结
                        </a>
                    </li>
                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-14336.html'>
                            绝命毒师第三季完结
                        </a>
                    </li>
                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-14338.html'>
                            绝命毒师第四季完结
                        </a>
                    </li>
                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-14339.html'>
                            绝命毒师第五季完结
                        </a>
                    </li>

                </ul>
            </div>
        </div>


    </div>
`)
	buffer.WriteString(`



      	<div class="mainleft" id="content">
			<div class="article_container row  box">
				<h1 class="info-name">`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["info"].(map[string]string)["name"], buffer)
	buffer.WriteString(`</h1>
                    <div class="article_info">
                        <span class="info_author info_ico"></span>
                        <span class="info_category info_ico"><a href="" title="查看剧情片中的全部文章" rel="category tag">`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["detail"].(map[string]interface{})["type"].(string), buffer)
	buffer.WriteString(`</a></span>
                        <span class="info_date info_ico">`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["detail"].(map[string]interface{})["update"].(string), buffer)
	buffer.WriteString(`</span>
                    </div>
            	<div class="clear"></div>

            <div class="context">
				<div id="post_content">
                    <p>
                        <img
                                src="/static/image/golang.png"
                                data-url="`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["info"].(map[string]string)["cover"], buffer)
	buffer.WriteString(`"
                                onerror="this.src='/static/image/golang.png'"
                                alt="`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["info"].(map[string]string)["name"], buffer)
	buffer.WriteString(`" >

                        <br>◎片　　名　`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["info"].(map[string]string)["name"], buffer)
	buffer.WriteString(`
                        <br>◎别　　名　`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["detail"].(map[string]interface{})["alias"].(string), buffer)
	buffer.WriteString(`
                        <br>◎年　　代　`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["detail"].(map[string]interface{})["released"].(string), buffer)
	buffer.WriteString(`
                        <br>◎产　　地　`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["detail"].(map[string]interface{})["area"].(string), buffer)
	buffer.WriteString(`
                        <br>◎类　　别　`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["detail"].(map[string]interface{})["type"].(string), buffer)
	buffer.WriteString(`
                        <br>◎语　　言　`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["detail"].(map[string]interface{})["language"].(string), buffer)
	buffer.WriteString(`
                        <br>◎片　　长　`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["detail"].(map[string]interface{})["length"].(string), buffer)
	buffer.WriteString(` 分钟
                        <br>◎导　　演　`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["detail"].(map[string]interface{})["director"].(string), buffer)
	buffer.WriteString(`
                        <br>◎主　　演　`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["detail"].(map[string]interface{})["starring"].(string), buffer)
	buffer.WriteString(`
                    </p>
                    <p>◎简　　介
                        <br>
                        `)
	if show["MovieDetail"].(map[string]interface{})["detail"].(map[string]interface{})["vod_play_info"] != nil {
		hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["detail"].(map[string]interface{})["vod_play_info"].(string), buffer)
	}
	buffer.WriteString(`
                        <br>&nbsp;
                    </p>

                    <hr>

</div>
                &nbsp;               	<div class="clear"></div>

<div class="widget box row">
    <h3>播放地址 kuyun （无插件 极速播放）</h3>

    `)
	if show["MovieDetail"].(map[string]interface{})["is_film"].(string) == "1" {
		for _, MovieDetail := range show["MovieDetail"].(map[string]interface{})["kuyun"].([]map[string]interface{}) {
			buffer.WriteString(`
    <a title='在线播放' href="javascript:;" play-link ="`)
			hero.EscapeHTML(MovieDetail["play_link"].(string), buffer)
			buffer.WriteString(`"  play-type="kuyun" class="lBtn">
        在线播放 `)
			hero.EscapeHTML(MovieDetail["episode"].(string), buffer)
			buffer.WriteString(`
    </a>

    `)
		}
	} else {
		for _, MovieDetail := range show["MovieDetail"].(map[string]interface{})["kuyun"].([]map[string]interface{}) {
			buffer.WriteString(`
    <a title='剧集' href="javascript:;"  play-link ="`)
			hero.EscapeHTML(MovieDetail["play_link"].(string), buffer)
			buffer.WriteString(`"  play-type="kuyun" class="lBtn">
        第`)
			hero.EscapeHTML(MovieDetail["episode"].(string), buffer)
			buffer.WriteString(`集
    </a>

    `)
		}
	}
	buffer.WriteString(`

</div>

<div class="widget box row">
    <h3>播放地址 ckm3u8（无插件 极速播放）</h3>
    `)
	if show["MovieDetail"].(map[string]interface{})["is_film"].(string) == "1" {
		for _, MovieDetail := range show["MovieDetail"].(map[string]interface{})["ckm3u8"].([]map[string]interface{}) {
			buffer.WriteString(`
        <a title='在线播放' href="javascript:;"  play-link ="`)
			hero.EscapeHTML(MovieDetail["play_link"].(string), buffer)
			buffer.WriteString(`"  play-type="m3u8" class="lBtn">
            在线播放 `)
			hero.EscapeHTML(MovieDetail["episode"].(string), buffer)
			buffer.WriteString(`
        </a>

    `)
		}
	} else {
		for _, MovieDetail := range show["MovieDetail"].(map[string]interface{})["ckm3u8"].([]map[string]interface{}) {
			buffer.WriteString(`
    <a title='剧集' href="javascript:;" play-link ="`)
			hero.EscapeHTML(MovieDetail["play_link"].(string), buffer)
			buffer.WriteString(`"  play-type="m3u8" class="lBtn">
        第`)
			hero.EscapeHTML(MovieDetail["episode"].(string), buffer)
			buffer.WriteString(`集
    </a>

    `)
		}
	}
	buffer.WriteString(`
</div>

<div class="widget box row">
    <h3>播放地址 MP4（无需安装插件，可下载）</h3>

    `)
	if show["MovieDetail"].(map[string]interface{})["is_film"].(string) == "1" {
		for _, MovieDetail := range show["MovieDetail"].(map[string]interface{})["download"].([]map[string]interface{}) {
			buffer.WriteString(`
    <a title='在线播放' href="javascript:;" play-link ="`)
			hero.EscapeHTML(MovieDetail["play_link"].(string), buffer)
			buffer.WriteString(`"  play-type="mp4" class="lBtn">
        在线播放 `)
			hero.EscapeHTML(MovieDetail["episode"].(string), buffer)
			buffer.WriteString(`
    </a>

    `)
		}
	} else {
		for _, MovieDetail := range show["MovieDetail"].(map[string]interface{})["download"].([]map[string]interface{}) {
			buffer.WriteString(`
    <a title='剧集' href="javascript:;"  play-link ="`)
			hero.EscapeHTML(MovieDetail["play_link"].(string), buffer)
			buffer.WriteString(`" play-type="mp4" class="lBtn">
        第`)
			hero.EscapeHTML(MovieDetail["episode"].(string), buffer)
			buffer.WriteString(`集
    </a>

    `)
		}
	}
	buffer.WriteString(`
</div>

            </div>
		</div>

    	<div>

	</div>

  </div>
    </div>




</body>

`)
	buffer.WriteString(`
    <div class="clear"></div>
    <div id="footer">
        <div class="footnav container">
            <ul id="footnav" class="menu">
            </ul>
        </div>
        <div class="footnav container">
            <ul id="friendlink" class="menu">
            </ul>
        </div>
        <div class="copyright">
            <p> Copyright &copy; 2019- `)
	hero.EscapeHTML(time.Now().Format("2006"), buffer)
	buffer.WriteString(`
                <a href="">
                    <strong> GoMovies </strong>
                </a>
            </p>

            <p> 本站所有视频均由程序自动采集而来，版权归原创者所有，如果侵犯了你的权益，请通知我删除侵权内容，谢谢合作。 </p>

            <p class="author"> power by
                <a href="https://hzz.cool" target="_blank" rel="external">
                    hzz.cool
                </a>
            </p>
        </div>
    </div>

    <!--gototop-->
    <div id="tbox"> <a id="gotop" href="javascript:void(0)"></a> </div>


    <script type='text/javascript' src='/static/js/jquery.min-3.8.1.js'></script>
    <script type='text/javascript' src='/static/js/lets-kill-ie6-3.8.1.js'></script>

    <script type="text/javascript" src="/static/js/jquery.masonry.js"></script>
    <script type="text/javascript" src="/static/js/loostrive.js"></script>

    <script language="javascript" type="text/javascript">
        (function() {
            var oDiv = document.getElementById("nav_b");
            var H = 0, iE6;
            var Y = oDiv;
            while (Y) {
                H += Y.offsetTop;
                Y = Y.offsetParent;
            };
            iE6 = window.ActiveXObject && !window.XMLHttpRequest;
            if (!iE6) {
                window.onscroll = function() {
                    var s = document.body.scrollTop || document.documentElement.scrollTop;
                    if (s > H) {
                        oDiv.className = "mainmenus container nav_b";
                        if (iE6) {
                            oDiv.style.top = (s - H) + "px";
                        }
                    } else {
                        oDiv.className = "mainmenus container";
                    }
                }
            }
        })();
    </script>


    <script>
        function replaceImg(){
            $("img").each(function () {
                let realImgUrl = $(this).attr("data-url");
                if ( realImgUrl !== "" )
                {
                    $(this).attr("src",$(this).attr("data-url"))
                }
            });
        }
        setTimeout(replaceImg, 1000);
    </script>

    <style>
        .search_m{display: none;}
        @media only screen and (max-width: 640px){
            .search_m{max-width: 360px !important;margin: 0 auto;margin-bottom: 10px;display: block;}
        }
    </style>
`)
	buffer.WriteString(`

<script>
    $('.lBtn').on('click', function(){
        let play_url = $(this).attr("play-link");
        let play_type = $(this).attr("play-type");

        let link = "`)
	hero.EscapeHTML(show["MovieDetail"].(map[string]interface{})["info"].(map[string]string)["link"], buffer)
	buffer.WriteString(`";

        window.open("//" +location.host + '/play?play_url=' + play_url + '&play_type=' + play_type + '&real_play=1'+'&link='+ link);
    });
</script>

</html>
`)

}
