(function(){

$('.close').click(function(){
	$(this.parentNode).remove();
});

$('.form-input').keydown(function(e){
	if ( event.which == 13 ) {
		$(".content-labels").append("<div class=\"content-label\">"+ $('.form-input').val() + "<span class=\"close\">x</span></div>");
		$('.form-input').val("");
	}
	$('.close').click(function(){
		$(this.parentNode).remove();
	});
});

})();