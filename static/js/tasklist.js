$(function() {
	  getTaskList();
  
    //任务列表
    function getTaskList(){
    		console.log(window.location)
        axios.get(window.location.origin+"/"+'tasklist').then(function(result){
            //console.log(result.data);
            rows = result.data.data;
			console.log("length==",rows.length);
			var tr = "";
            for(var i = 0; i < rows.length; i++){
                

                tr += '<tr>';
                tr += '<td class="active">'+rows[i].task_id+'</td>';
                tr += '<td class="success">'+rows[i].user_name+'</td>';
                tr += '<td class="warning">'+rows[i].task_user+'</td>';
                tr += '<td class="danger">'+rows[i].bonus+'</td>';
                tr += '<td class="info">'+rows[i].status+'</td>';
                tr += '<td class="active">'+rows[i].remark+'</td>';
                tr += '<td class="success">'+rows[i].comment+'</td>';
                tr += '</tr>';
                
            }
            $("#list").html(tr);
            
        })
        .catch(error => console.log(error));    
    }
    
});