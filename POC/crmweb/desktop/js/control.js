function SaveSpkdAddWSController($scope, $http) {
	$scope.errors = [];
	$scope.msgs = [];

//	$scope.USER_CODE= 'LLTHUNYADAP';
//	$scope.BLPD_INDC= 'PCN';
//	$scope.CS_SPKD_PCN__CUST_NUMB= '536672462';
//	$scope.CS_SPKD_PCN__SUBR_NUMB= '66900010040';
//	$scope.CS_SPKD_PCN__PACK_CODE= '31001501';
//	$scope.RD_TELP__TELP_TYPE= 'TEL';
//	$scope.SAVE_FLAG= '1';

	$scope.Save= function() {

		$scope.errors.splice(0, $scope.errors.length); // remove all error messages
		$scope.msgs.splice(0, $scope.msgs.length);

		$http({
			url: '/post',
			method: 'POST',
			headers: {
			  'Accept': 'application/json, text/javascript', 
			  'Content-Type': 'application/json; charset=utf-8'
			},
			data: {
				'USER_CODE': $scope.USER_CODE, 
				'BLPD_INDC': $scope.BLPD_INDC, 
				'CS_SPKD_PCN__CUST_NUMB': $scope.CS_SPKD_PCN__CUST_NUMB,
				'CS_SPKD_PCN__SUBR_NUMB': $scope.CS_SPKD_PCN__SUBR_NUMB,
				'CS_SPKD_PCN__PACK_CODE': $scope.CS_SPKD_PCN__PACK_CODE,
				'CS_SUBR_PCN__SUBR_TYPE': $scope.CS_SUBR_PCN__SUBR_TYPE,
				'RD_TELP__TELP_TYPE': $scope.RD_TELP__TELP_TYPE,
				'SAVE_FLAG': $scope.SAVE_FLAG
			},
		}).success(function(data){ 
			$scope.Status = data.Status;
			if(data.Status == "success"){
//			$scope.CS_SPKD_PCN__PACK_CODE=data.CS_SPKD_PCN__PACK_CODE;
//			$scope.CS_SPKD_PCN__PACK_CODE=data.CS_SPKD_PCN__PACK_CODE;
//			$scope.CS_PKPL_PCN__PACK_DESC=data.CS_PKPL_PCN__PACK_DESC;
//			$scope.CS_PACK_TYPE__PACK_TYPE_DESC=data.CS_PACK_TYPE__PACK_TYPE_DESC;
//			$scope.CS_SPKD_PCN__PACK_STRT_DTTM=data.CS_SPKD_PCN__PACK_STRT_DTTM;
//			$scope.CS_SPKD_PCN__PACK_END_DTTM=data.CS_SPKD_PCN__PACK_END_DTTM;
//			$scope.CS_SPKD_PCN__DISC_CODE=data.CS_SPKD_PCN__DISC_CODE;
//			$scope.TBL_OCCR=data.TBL_OCCR;
			}else{
//				$scope.TBL_OCCR=data.Faultcode+data.Faultstring;
			}

//			if (data.Msg != '')
//			{
//				$scope.msgs.push(data.SPKD_PCN__PACK_CODE);
//			}
//			else
//			{
//				$scope.errors.push(data.SPKD_PCN__PACK_CODE);
//			}
				$scope.USER_CODE= '';
				$scope.BLPD_INDC= '';
				$scope.CS_SPKD_PCN__CUST_NUMB= '';
				$scope.CS_SPKD_PCN__SUBR_NUMB= '';
				$scope.CS_SPKD_PCN__PACK_CODE= '';
				$scope.RD_TELP__TELP_TYPE= '';
				$scope.SAVE_FLAG= '';			
		}).error(function(data, status) { // called asynchronously if an error occurs
// or server returns response with an error status.
			$scope.errors.push("error");
		});
		
	}
}