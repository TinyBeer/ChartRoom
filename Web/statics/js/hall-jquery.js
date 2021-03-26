
$('#btn1').click(function(){
  axios.post('content', {content:$("#input_content").val()})
  .then(function (response) {
    console.log(response) 
    $('#input_content').val('')                         
  })
  .catch(function (error) {
    console.log(error);
  });
})


$(function(){
    
    var timer = setInterval("clock()", 2000)
    
    console.log("Width:"+window.innerWidth);
    console.log("Height:"+window.innerHeight);
    if (window.innerWidth / window.innerHeight < 1.5){
      $("#dialog-container").css({
        "width" : "80%",
        "height" : "80%"
    })
    }
}); 

function clock() {
  // NewTalk('hello');
  axios.get('content', {})
        .then(function (response) {
        if (response.data.content.length != 0){
            NewTalk(response.data.content)
          }                            
        })
        .catch(function (error) {
          console.log(error);
        });
}


function NewTalk(content) {
        var $liNew = $(' <li><span>'+ content + '</span></li>');
        $liNew.appendTo('#ul1');
        $liNew.hide();
        $liNew.slideDown(1000);
        
        removeFirst();
    }

function removeFirst(){
        if ($("#ul1").children("li:last-child").offset().top  >
        $("#input-container").offset().top -80  ){
                $('#ul1').children("li").first().slideUp(1000, function(){
                    $(this).remove();
                    removeFirst();
         })
    }
            
}