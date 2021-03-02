# Android and Javascript by WebView

[[How-to][referral]]:
>1. Android calls JS
1. `WebView.loadUrl()`
1. `WebView.evaluateJavascript()`

>2. JS calls Android
1. `WebView.addJavaScriptInterface()`
1. `WebViewClient.shouldOverrideUrlLoading()` intercepting URL changes
1. `WebChromClient`, `onJsAlert()`,`onJsConfirm()`, `onJsPrompt()` intercepting `alert()`, `confirm()` and `prompt()`

<br>

---
## Android Calls JS
1. `WebView.loadUrl()`
1. `WebView.evaluateJavascript()`

Compare:

Method | Pros | Cons | Senario
---|---|---|---
`loadUrl` | | hassle of returning results | No need for results
`evaluateJavascript` | efficient | 4.4 above | > 4.4

`WebView.loadUrl()` -> js file url -> `js execute function`

>Local file `file:///android_asset/js.html`

```java
WebSettings settings = WebView.getSettings();
settings.setJavaScriptEnabled(true);
settings.setJavaScriptCanOpenWindowsAutomatically(true);

// load js file first
webView.loadUrl(LOCAL_FILE | Remote_Address)

// those javascript page only available after page fully loaded
button.setOnClickListener(new View.OnClickListener() {
    // ...
    webView.loadUrl("javascript:js_function()")
})

// catch js alert and pop as Android.AlertDialog
mWebView.setWebChromeClient(new WebChromeClient() {
    @Override
    public boolean onJsAlert(WebView view, String url, String message, final JsResult result) {
        AlertDialog.Builder b = new AlertDialog.Builder(MainActivity.this);
        b.setTitle("Alert");
        b.setMessage(message);
        b.setPositiveButton(android.R.string.ok, new DialogInterface.OnClickListener() {
            @Override
            public void onClick(DialogInterface dialog, int which) {
                result.confirm();
            }
        });
        b.setCancelable(false);
        b.create().show();
        return true;
    }

});
```

<br>

### `evaluateJavascript`
```java
webView.evaluateJavascript("javascript:call_js_function()", new ValueCallback<String>() {
    @Override
    public void onReceiveValue(String value) {}
})
```
<br>
<br>

---
## 2. JS calls Android
1. `WebView.addJavaScriptInterface()`
1. `WebViewClient.shouldOverrideUrlLoading()` intercepting URL changes
1. `WebChromClient`, `onJsAlert()`,`onJsConfirm()`, `onJsPrompt()` intercepting `alert()`, `confirm()` and `prompt()`

Compare:

Method | Pros | Cons | Senario
---|---|---|---
`addJavascriptInterface` | mapping object into JS | > 4.2 | > 4.2
`shouldOverrideUrlLoading` | good | protocols | iOS method
`onJsPrompt` | good | protocols | mostly usable

### 1. Map Android Object into JS
`WebView.addJavascriptInterface(new ClassName(), "jsname")` will inject an object with `@JavascriptInterface` method exposed into javascript

### 2. Intercept URL Change
1. `JS location` changes
1. `WebView.shouldOverrideUrlLoading`
1. parse URL and behave differently
1. not easy to object results

```javascript
// callback for results
function onResults(value) {
    // will be called by WebView here
}

document.location = "js://webview?arg1=111&arg2=222";
```

```java
webView.setWebViewClient(new WebViewClient() {
    @Override
    public boolean shouldOverrideUrlLoading(WebView view, String url) {
        // may judge by URL scheme and authority
        Uri uri = Uri.parse(url);
        if ("js".equals(uri.getScheme()) && "webview".equals(uri.getAuthority())) {
            Set<String> paramNames = uri.getQueryParameterNames()
            new Thread(new Runnable() {
                public void run() {
                    // do stuff
                    // notify results
                    webView.evaluateJavascript("onResult(" + result + ")")
                }
            }).start()

            // stop here
            return true;
        }
        return super.shouldOverrideUrlLoading(view, url);
    }
})
```
<br>

### 3. Intercept `alert`, `confirm` and `prompt`
1. JS alert, confirm or promts with data
1. `WebChromeClient` intercept it
1. parse passed data
1. comparatively easier to get results

```javascript
var result = prompt("js://demo?arg1=111&arg2=222");
alert("result" + result)
```

```java
webView.setWebChromeClient(new WebChromeClient() {
    public boolean onJsPrompt(WebView view, String url, String message, String defaultValue, JsPromptResult result) {
        // parse payload like url
        Uri uri = Uri.parse(message);
        if ("js".equals(uri.getScheme()) && "demo".equals(uri.getAuthority())) {
            // pass value back to JS prompt
            result.confirm("result");
        }
        // stop here
        return true;
    }
    return super.onJsPrompt(...);
})

```

[referral]: https://blog.csdn.net/carson_ho/article/details/64904691