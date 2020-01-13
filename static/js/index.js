$(function(){


  $('#entry_btn').on('click',function(){
	var username=$("#user").val();
    var password=$("#password").val();
    console.log(window.location)
    axios.post(window.location.href+"login",{
    		user: username,
		pass: password
    }).then((response)=>{
    		console.log(response)
    		respmsg = response.data
		if (respmsg.code == "0") {
			console.log("login sucess")
			window.open(window.location.origin+"/"+"tasklist.html", "_self")
		} else {
			console.log("login failed")
		}
		
    }).catch((err)=>{
    		console.log(err)
    })
  })
})