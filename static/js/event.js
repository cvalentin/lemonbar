(function(){

	var resetVisibility = function(){
		$('.content-hover').css('visibility', 'hidden');
		$('.content-add').css('visibility', 'visible');
	};

	var addGameToEvent = function(){
		$('.event-grid').append('<div class=\"game-object\"><div class=\"game-object-title\">Game Title</div><div class="game-object-timer">00:00</div></div>');
		var games = $('.game-object');
		startTimer($(games)[$(games).length - 1]);
	};

	var startTimer = function(game){
		var timer = $(game).children();
		var timerUpdate = $(timer[timer.length - 1]);
		
		var start = new Date().getTime(),
	    	time = 0,
	    	elapsed = '00:00',
	    	timerEnd;

		function instance()
		{
		    time += 1000;

		    elapsed = Math.floor(time / 100) / 10;
		    var leading = '';
		    var minute = Math.floor(elapsed / 60);
		    if(Math.round(elapsed) == elapsed) {
		    	if(elapsed % 60 < 10){
		    		leading = '0';
		    	}
		    	if(minute % 60 < 10){
		    		minute = '0' + minute;
		    	}
		    }

		    timerUpdate.html(minute + ':' + leading + (elapsed % 60));

		    var diff = (new Date().getTime() - start) - time;
		    timerEnd = window.setTimeout(instance, (1000 - diff));
		}

		timerEnd = window.setTimeout(instance, 1000);

		$(game).append('<div class=\"button game-state-end\">End</div>');
		var endButton = $(game).find('.game-state-end')[0];
		$(endButton).click(function(){
			console.log('click');
			clearInterval(timerEnd);
		});
	};

	$('.content-add').click(function(){
		$('.content-hover').css('visibility', 'visible');
		$('.content-add').css('visibility', 'hidden');
	});

	$('.game-state-begin').click(function(){
		resetVisibility();
		addGameToEvent();
	});

	$('.content-form-close').click(function(){
		resetVisibility();
	});

})();