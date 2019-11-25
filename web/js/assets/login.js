
$(function(){
	$(document).on('click', '#btn-submit', function(){
		var form = $('#form')[0];
		form.action = "/top";
		form.submit();
	});
});
