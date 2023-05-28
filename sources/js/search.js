$("#search-icon").on('click', function () {
    var display = $("#search-div").css("display");
    if(display =="none"){
        $("#search-div").css('display', 'block');
        $(".wc").css('opacity', '20%');
    }else{
        $("#search-div").css('display', 'none');
        $(".wc").css('opacity', '100%');
    }
    
})