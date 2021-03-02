# Android Hybrid Core: `WebView`
[[how-to][referral]]  
[[WebView vulnerables][WebView-vulnerables]]

`Chrome` engine after Version 4.4

What it can do:
1. render web pages
1. load either local or remote resources
1. interact with javascript

Related classes:
* `WebView`
* `WebSettings`
* `WebViewClient` deals with page events
* `WebChromeClient` deals with javascript interactions

---

## `WebView` 

Method | Purpose
---|---
onResume() | activate WebView
onPause() | stop kernel functions like DOM parse, plugin invocation, javascript execution
pauseTimers() | 
resumeTimers() | 
layout.removeView(webView) <br> webView.destroy() | completely remove webView

<br>
<br>

### 0. Create a `WebView`
Permission
`<uses-permission android:name="android.permission.INTERNET"/>`

```java
// from current Activity
WebView webView = new WebView(this);
// from layout resource 
WebView webView = (WebView) findViewById(R.id.webViewId);
```

### 1. Forward/Backward Softkey
Method | Purpose
---|---
canGoBack() |
goBack() |
canGoForward() |
goForward() |
goBackOrForward(nSteps) |

>Disable default **`BACK`** function

By default, system **`BACK`** will exit not go back to previous page. In order to do so, disable `back` event:
```java
public boolean onKeyDown(int keyCod, KeyVent event) {
    if ((keyCode == KEYCODE_BACK) && webView.canGoBack()) {
        // go back page
        webView.goBack();
        return true;
    }
    return super.onKeyDown(keyCode, event);
}
```
<br>

### 2. Cache
Method | Purpose
---|---
clearCache(true) | 
clearHistory() | 
clearFormData() | 

---
<br>
<br>

 
## `WebViewSettings`
```java
WebSettings settings = webView.getSettings();

// interact with javascript
settings.setJavaScriptEnabled(true);

// enable plugin
settings.setPluginsEnabled(true);

// on screen RWD
settings.setUseWideViewPort(true);
settings.setLoadWithOverviewMode(true);

// webview's zoom control
settings.setSupportZoom(true);
settings.setBuiltInZoomControls(true);
settings.setDisplayZoomControls(false);

// other
// disable cache
settings.setCacheMode(WebSettings.LOAD_CACHE_ELSE_NETWORK);
settings.setAllowFileAccess(true);
settings.setJavascriptCanOpenWindowsAutomatically(true);
settings.setLoadsImageAutomatically(true);
settings.setDefaultTextEncodingName("utr-8");
```
>Cache
1. when loading pages, in `/data/data/` creates `database` and `cache` folders
1. requested URL recorded in `WebViewCache.db`, and page cache files in `WebViewCache` folder
1. strategy
`webView.getSettings().setCacheMode (mode)`

Mode | Purpose
---|---
LOAD_CACHE_ONLY | no remote, only local
LOAD_DEFAULT | refer to `cache-control`
LOAD_NO_CACHE | only from remote
LOAD_CACHE_ELSE_NETWORK | local first then remote

4. offline
```java
if (NetStatusUtil.isConnected(getApplicationContext())) {
    webSettings.setCacheMode(WebSettings.LOAD_DEFAULT); // refer to cache-control
} else {
    webSettings.setCacheMode(WebSettings.LOAD_CACHE_ELSE_NETWORK);
}

// enable DOM storage API
settings.setDomStorageEnabled(true);
// enable database storage API
settings.setDatabaseEnable(true);

// enable cache and set cache file path
settings.setAppCacheEnabled(true);
settings.setAppCachePath(
    getFilesDir().getAbsolutPath() + APP_CACHE_DIRNAME
)
```
---
## `WebViewClient`
Method | Purpose
---|---
shouldOverrideUrlLoading() | intercept before changing `javascript.window.location.href`
onPageStarted() | set loading page ???
onPageFinished() |
onLoadResource() | called once per url resource
onReceivedError() | called when status code rather than 2xx
onReceivedSslError() |

URL Kinds | Purpose
---|---
https://www.google.com/ | normal remote page
file:///android_asset/test.html | load page from apk
content://com.android.htmlfileprovider/sdcard/test.html | load phone local page

---
## `WebChromeClient`
Method | Purpose
---|---
onProgressChanged() | page loading progress
onReceivedTitle() | intercept page title

<br>

---
## Memory Leakage with `WebView`
>1. instantiate not declare
```java
LinearLayout.LayoutParams params = new LinearLayout.LayoutParams(ViewGroup.LayoutParams.MATCH_PARENT, ViewGroup.LayoutParams.MATCH_PARENT);
        mWebView = new WebView(getApplicationContext());
mWebView.setLayoutParams(params);
mLayout.addView(mWebView);
``` 
>2. destroy
1. load null content
1. clear history
1. remove from layout
1. destroy
1. set null
```java
@Override
protected void onDestroy() {
    if (mWebView != null) {
        mWebView.loadDataWithBaseURL(null, "", "text/html", "utf-8", null);
        mWebView.clearHistory();

        ((ViewGroup) mWebView.getParent()).removeView(mWebView);
        mWebView.destroy();
        mWebView = null;
    }
    super.onDestroy();
}
```

[referral]: https://blog.csdn.net/carson_ho/article/details/52693322
[WebView-vulnerables]: https://www.jianshu.com/p/3a345d27cd42