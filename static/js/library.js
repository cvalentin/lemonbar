(function(){

$('.close').click(function(){
	$(this.parentNode).remove();
});

$('.form-input').keydown(function(e){
	var input = $('.form-input').val();
	if ( event.which == 13 && 1 <= input.length) {
		$('.content-labels').append('<div class=\'content-label\'>'+ input + '<span class=\'close\'>x</span></div>');
		$('.form-input').val('');
	}
	$('.close').click(function(){
		$(this.parentNode).remove();
	});
});

})();