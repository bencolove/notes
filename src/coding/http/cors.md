# Cors 

## Spring MVC

Like [[this post][OPTIONS-preflight]] describe the Cors issue:
> Request header field Authorization is not allowed by Access-Control-Allow-Headers in preflight response.

`addCorsMapping` can configure spring-managed Cors to specific routes. It is alright for any http client rather then a browser like Chrome.

Because **any custom header** like custom `Authorization` will cause a preflight `OPTIONS` before the real deal and it
requires `Content-Type` and all custom headers to be allowed, and `*` just simply doesnot work !!!

So a dedicated _`HandlerInteceptor`_ should be adopted to handle `OPTIONS` preflight exclusively.

```java
public class FodWebConfig extends WebMvcConfigurationSupport {

    @Override
    protected void addCorsMappings(CorsRegistry registry) {
        registry.addMapping("/fod/**")
            .allowedOrigins("*")
            .allowedHeaders("*")
            .allowedMethods("*");
    }

     @Override
    protected void addInterceptors(InterceptorRegistry registry) {

        // 放行OPTIONS preflighted 請求
        registry.addInterceptor(new CorsInterceptor())
                .addPathPatterns("/fod/**");
    }
    ...
}

public class CorsInterceptor extends HandlerInterceptorAdapter {

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        if (RequestMethod.OPTIONS.name().equalsIgnoreCase(request.getMethod())) {
            // since Authorization header is used to present token, therefore preflight is applied
            // OPTIONS preflighted, let it through
            response.setHeader("Access-Control-Allow-Origin", "*");
            response.setHeader("Access-Control-Allow-Methods", "*");
            response.setHeader("Access-Control-Allow-Headers", "Authorization, Content-Type");
            response.setHeader("Access-Control-Max-Age", "86400");
            response.setStatus(HttpStatus.NO_CONTENT.value());
            return false;
        }
        return super.preHandle(request, response, handler);
    }
}
```

[OPTIONS-preflight]: https://zhuanlan.zhihu.com/p/42847303