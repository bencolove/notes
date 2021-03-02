# The App
[[guide][guide]]

> Declare app components in manifest:
* activity
* fragment
* service
* content provider
* brocast receiver

> Design Principles
1. Deparation of concerns
    * view displays data (without auto binding)
    * logics in ViewModel
1. Data-driven UI (persistent, how?)
    * bind view and data
    * only update data
    * let data update view  

>What they look like

![](../res/arch/final-architecture.png)


>Simple Project Layout
1. user_profile.xml -- UI layout
1. UserProfileFrament.java -- UI controller displaying data
1. UserProfileViewModel.java -- prepares data to UI and reacts to user interactions
1. UserRepository -- abstract data source
1. UserApi -- remote Api to fetch data
1. ORM Room
1. DI Hilt

--_`UserProfileViewModel`_--

```java
class UserProfileViewModel @ViewModelInject constructor(
    @Assisted savedState: SavedStateHandle,
    userRepository: UserRepository
) : ViewModel() {
    // get it back from saved state
   val userId : String = savedState["uid"]?: throw IllegalArgumentException("missing uid")
   // use live data to recieve updates
   private val _user = MutableLiveData<User>()
   val user = LiveData<User> = _user

   init {
       // where this scope comes from
       viewModelScope.launch {
           // fetch and notify
           _user.value = userRepository.getUser(userId)
       }
   }
}
```

--_`UserProfileFragment`_--

```java
class UserProfileFragment : Fragment() {
   // To use the viewModels() extension function, include
   // "androidx.fragment:fragment-ktx:latest-version" in your app
   // module's build.gradle file.
   private val viewModel: UserProfileViewModel by viewModels(
       factoryProducer = {
           // save state for resume
           SavedStateVMFactory(this)
       }
   )

   override fun onCreateView(
       inflater: LayoutInflater, container: ViewGroup?,
       savedInstanceState: Bundle?
   ): View {
       return inflater.inflate(R.layout.main_fragment, container, false)
   }

   override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
       super.onViewCreated(view, savedInstanceState)
       // connect to view model's live data
       // removed when the lifecycle destroyed
       viewModel.user.observe(viewLifecycleOwner) {
           // update UI due to changes
       }
   }
}
```

--_`UserApi`_--
```java
interface Webservice {
   @GET("/users/{user}")
   suspend fun getUser(@Path("user") userId: String): User
}
```

--_`UserRepository`_--

```java
// @Inject by Hilt
class UserRepository @Inject constructor (
   private val webservice: Webservice,
   private val userCache: UserCache
   // ...
   suspend fun getUser(userId: String) =
       // should cache
       val cached: User? = userCache.get(userId)
       if (cached != null) {
           return cached
       }
       val freshUser = webservice.getUser(userId)
       userCahce.put(userId, freshUser)
       return freshUser
}
```

[guide]: https://developer.android.com/jetpack/guide