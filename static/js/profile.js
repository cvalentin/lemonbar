(function(){

	var games = [];

	var resetVisibility = function () {
		$('.content-hover').css('visibility', 'hidden');
		$('.content-add').css('visibility', 'visible');
	};

	var resetContent = function () {
		games = [];
		$('.content-label').remove();
	};

	$('.content-add').click(function () {
		$('.content-hover').css('visibility', 'visible');
		$('.form-input').focus();
		$('.content-add').css('visibility', 'hidden');
	});

	$('.content-form-close').click(function () {
		resetVisibility();
		resetContent();
	});

	$('.close').click(function () {
		$(this.parentNode).remove();
	});

	$('.form-input').keydown(function (e) {
		var input = $('.form-input').val();
		if ( event.which == 13 && 1 <= input.length) {
			$('.content-labels').append('<div class=\'content-label\'>'+ input + '<span class=\'close\'>x</span></div>');
			games.push({
				userID:'temp',
				game:input
			});
			$('.form-input').val('');
		}
		$('.close').click(function () {
			$(this.parentNode).remove();
		});
	});

	$('.library-game-add').click(function () {
		console.log(games);
		//push data to ...wherever

		resetVisibility();
		resetContent();
	});

})();