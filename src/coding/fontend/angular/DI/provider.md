# Dependency Provider

    A dependency provider configures an injector with a DI token used to decide the runtime value.

> Concepts:
* injector -- runtime to inject values where asked
* provider -- configure how to find values

> Where is `providers` defined:

`@NgModule()` on `AppModule` or `@Components`

> How injector manage its values for provision

Injector utilizes sth like map with key of type `DI token` (simply `class`, can not be value of `interface`, nor string literal) to lookup its provider which configures how to find/instantiate/alias the injecting value.

> Define providers

`providers: [Logger]` 
is expanded to  `[{ provide: Logger, useClass: Logger }]`
* `provide`: token to lookup in injector's DI map 
* `useClass`: new 
* `useExisting`: alias DI token
* `useValue`:
* `useFactory`: 

> 1. provider `useClass`


>2. provider `useExisting`

1. alias already exposed DI token  

`providers: [{provide: OldLogger, useExisting: NewLogger}]`

2. alias class (no corresponding DI token defined)

`[{provide: OldLogger, useExisting: forwardRef(() => NewLoggerClass)}]`

with helper

>3. provider `useValue` injecting objects

1. inject plain object (duck typing)
1. inject configuration object

Non-class like typescript `interfaces` will be erased after transpiled into javascript, so `interfaces` can not be DI tokens nor injected.

It has to be with the help of `InjectionToken`:

```typescript
import { InjectionToken } from '@angular/core';

export const APP_CONFIG = new InjectionToken<AppConfig>('app.config');

export const MY_CONFIG: AppConfig = {
    endpoint: 'api.hero.com',
    title: 'Non-class DI Token'
}

// providing
providers: [{
    provide: APP_CONFIG, useValue: MY_CONFIG
}]
// injecting
constructor(@Inject(APP_CONFIG) config: AppConfig)
```

>4. provider `useFactory` dynamic provision

```typescript
// factory function, => HeroService
const heroServiceFactory = (logger: Logger, userService: UserService) => {
    return new HeroService(logger, userService.user)
}

// define provier
export let heroServiceProvider = {
    provide: HeroService,
    useFactory: heroServiceFactory,
    deps: [Logger, UserService]
};

// apply provider
import { heroServiceProvider } from './path/to';

@Component({
    selector: '',
    providers: [ heroServiceProvider ],
    template: ``
})

```