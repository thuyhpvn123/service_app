# Metanode-desktop-app service các hàm xử lý core

> databasePath = [docsDir stringByAppendingPathComponent: @".meta_browser.db"]];

## CheckNameWalletIsExist

> _input: name_

> querySQL int[]

> SELECT \* FROM walletTB WHERE name = **_name from input_**

```
    if (sqlite3_open(dbpath, &database) == SQLITE_OK) {
        while (sqlite3_step(statement) == SQLITE_ROW) {
            NSString *identity = [[NSString alloc] initWithUTF8String(const char *) sqlite3_column_text(statement, 0)];
            NSString *address = [[NSString alloc] initWithUTF8String(const char *) sqlite3_column_text(statement, 1)];
            NSString *name = [[NSString alloc] initWithUTF8String:(constchar *) sqlite3_column_text(statement, 2)];
            NSString *pendingBalance = [[NSString alloc]initWithUTF8String:(const char *) sqlite3_column_tex(statement, 3)];
            NSString *balance = [[NSString alloc] initWithUTF8String(const char *) sqlite3_column_text(statement, 4)];

            NSString * pendingBalanceHex = [Utils convertBigEdianToString[Utilities dataFromHexString:pendingBalance]];
            NSString * balanceHex =  [Utils convertBigEdianToString[Utilities dataFromHexString:balance]];

             wallet = {
                "id":identity,
                "address":address,
                "name": name,
                "pendingBalance": pendingBalance,
                "balance": balance,
                "pendingBalanceString": pendingBalanceHex,
                "balanceString": balanceHex,
            };
            resultArray.push(wallet);
        }
    }
    if ([resultArray count] > 0) {
        return resultArray length;
    }
    return 0;
```

## randomNameWallet (lastRandomString, timeCall)

> listData []

- Sirius
- Canopus
- Arcturus
- Alpha Centauri A
- Vega
- Rigel
- Procyon
- Achernar
- Betelgeuse
- Hadar (Agena)
- Capella A
- Altair
- Aldebaran
- Capella B
- Spica
- Antares
- Pollux
- Fomalhaut
- Deneb
- Mimosa

```
 int i = 0;
 int size = total của listData;
 string name = ""
 string lastRandom = lastRandomString
```

##### Tạo 1 vòng lập

`while (i < size){`

```
        uint32_t index = random của size;
        string randomValue = lấy 1 phần tử trong mảng listData random dựa vào index ở trên
        if (timeCall > 0) {
            randomValue = randomValue + timeCall
        }
        if (randomValue === lastRandom) {
            continue;
        }
        int isExist =
```

[checkNameWalletIsExist(randomValue)](#checknamewalletisexist)

```
    if (isExist == 0) {
        name = randomValue;
        break;
    }
    lastRandom = randomValue;
```

`   }`

##### kiểm tra name hoặc tiếp tục random 1 name mới

```
    if (name length > 0) {
        return name;
    }
    return randomNameWallet (timeCall + 1, lastRandom]) ;
```

## updateWalletUI

```
func updateWalletUI(data map[string]string) bool {
	databasePath := data["databasePath"]
	dbpath := []byte(databasePath)

	if sqlite3.Open(dbpath, &database) == sqlite3.SQLITE_OK {
		updateSQL := fmt.Sprintf("UPDATE walletTB SET bg = \"%s\", color = \"%s\", idSymbol = %d, pattern = \"%s\" WHERE address = \"%s\"", data["bg"], data["bg"], data["bg"], data["pattern"], data["address"])
		updateStmt := []byte(updateSQL)
		sqlite3.PrepareV2(database, updateStmt, -1, &statement, nil)
		sqlite3.Reset(statement)

		if sqlite3.Step(statement) == sqlite3.SQLITE_DONE {
			return true
		} else {
			return false
		}
	}
	return false
}
```

## updateBalance

```
func updateBalance(address string, pendingBalance string, balance string) bool {
	databasePath := data["databasePath"]
	dbpath := []byte(databasePath)

	if sqlite3.Open(dbpath, &database) == sqlite3.SQLITE_OK {
		bBalance := []byte(Utilities.DataFromHexString(balance))
		u256Balance := eevm.FromBigEndian(bBalance, 32)

		bPendingBalance := []byte(Utilities.DataFromHexString(pendingBalance))
		u256PendingBalance := eevm.FromBigEndian(bPendingBalance, 32)

		total := u256Balance + u256PendingBalance
		totalBalance := []byte(to_string(total))

		updateSQL := fmt.Sprintf("UPDATE walletTB SET pendingBalance = \"%s\", balance = \"%s\", totalBalance = \"%s\" WHERE address = \"%s\"", pendingBalance, balance, totalBalance, address)
		updateStmt := []byte(updateSQL)
		sqlite3.PrepareV2(database, updateStmt, -1, &statement, nil)
		sqlite3.Reset(statement)

		if sqlite3.Step(statement) == sqlite3.SQLITE_DONE {
			return true
		} else {
			return false
		}
	}
	return false
}

```

## getWalletByAddress

```
func getWalletByAddress(address string) []map[string]string {
	databasePath := data["databasePath"]
	dbpath := []byte(databasePath)

	if sqlite3.Open(dbpath, &database) == sqlite3.SQLITE_OK {
		querySQL := fmt.Sprintf("SELECT * FROM walletTB WHERE address = \"%s\"", address)
		queryStmt := []byte(querySQL)
		resultArray := make([]map[string]string, 0)

		if sqlite3.PrepareV2(database, queryStmt, -1, &statement, nil) == sqlite3.SQLITE_OK {
			for sqlite3.Step(statement) == sqlite3.SQLITE_ROW {
				walletFromStatement := getDataWalletFromStatement(statement)
				resultArray = append(resultArray, walletFromStatement)
			}
			sqlite3.Reset(statement)
			return resultArray
		}
	}
	sqlite3.Reset(statement)
	return nil
}

```

## insertTransaction

```
func insertTransaction(data map[string]string) bool {
	databasePath := data["databasePath"]
	dbpath := []byte(databasePath)

	if sqlite3.Open(dbpath, &database) == sqlite3.SQLITE_OK {
		hash := data["hash"]
		transactionDetail := getTransactionByHash(hash)

		if transactionDetail != nil {
			return true
		}

		type := data["type"]
		amount := data["amount"]
		fee := data["fee"]
		tip := data["tip"]

		for len(amount) < 64 {
			amount = "0" + amount
		}
		for len(fee) < 64 {
			fee = "0" + fee
		}
		for len(tip) < 64 {
			tip = "0" + tip
		}

		bAmount := []byte(Utilities.DataFromHexString(amount))
		u256Amount := eevm.FromBigEndian(bAmount, 32)

		bFee := []byte(Utilities.DataFromHexString(fee))
		u256Fee := eevm.FromBigEndian(bFee, 32)

		bTip := []byte(Utilities.DataFromHexString(tip))
		u256Tip := eevm.FromBigEndian(bTip, 32)

		total := u256Amount
		if u256Tip > 0 {
			total += u256Tip
		}
		if type == "send" {
			total += u256Fee
		}

		totalBalance := []byte(to_string(total))

		insertSQL := INSERT INTO transactionTB(hash, address, toAddress, pubKey, amount, pendingUse, balance, fee, tip, message, time, status, type, prevHash, sign, receive_info, isDeploy, isCall, functionCall, data, totalBalance, lastDeviceKey) VALUES(\"%@\",\"%@\",\"%@\",\"%@\",\"%@\",\"%@\",\"%@\",\"%@\",\"%@\",\"%@\",%ld,%ld,\"%@\",\"%@\",\"%@\",\"%@\",%ld,%ld,\"%@\",\"%@\",\"%@\",\"%@\")", data[@"hash"], data[@"address"], data[@"toAddress"], data[@"pubKey"], data[@"amount"], data[@"pendingUse"], data[@"balance"], data[@"fee"], data[@"tip"], data[@"message"], [data[@"time"] integerValue], [data[@"status"] integerValue], data[@"type"], data[@"prevHash"], data[@"sign"], data[@"receive_info"], [data[@"isDeploy"] integerValue], [data[@"isCall"] integerValue], data[@"functionCall"], data[@"data"], totalBalance, data[@"lastDeviceKey"] != NULL ? data[@"lastDeviceKey"] : @""];

        sqlite3_prepare_v2(database, insert_stmt,-1, &statement, NULL);

        sqlite3_reset(statement);

        if (sqlite3_step(statement) == SQLITE_DONE) {
            return YES;
        } else {
            return NO;
        }

```

> _output_ return YES or NO

## deleteAllTrans

> query

```
    DELETE FROM transactionTB
```

> _output_ return YES or NO

## getTotalTransactionSuccess (address)

```
    SELECT COUNT(*) FROM transactionTB WHERE type = \"send\" and status = 2 and address = \"%@\"",address
```

> _output_ return count;

## updateStatusTransactionByHash(status, hash)

Hàm sử dụng để cập nhật lại trạng thái của 1 transaction dựa vào

> Input:
>
> - status: Nhận vào giá trị của status muốn set cho transaction
> - hash: Mã hash của transaction muốn thay đổi status

Các lệnh query:

```
   @"SELECT max(time) FROM transactionTB where hash = \"%@\" ", hash
```

```
   @"UPDATE transactionTB SET status = %d where hash = \"%@\" AND time = %d "
```

> _Output_: return YES (nếu cập nhật status thành công) or NO(nếu cập nhật status thất bại hoặc xảy ra lỗi trong quá trình cập nhật);

## updateTransactionWithStatusPending(time)

Hàm sử dụng để cập nhật trạng thái của tất cả transaction thành fail khi thời gian khởi tạo của transaction không nhỏ hơn thời gian truyền vào và trạng thái của transaction là chưa thành công.

> Input:
>
> - time: Thời gian được truyền vào với kiểu int

Các lệnh query:

```
	@"UPDATE transactionTB SET status = 4 WHERE time >= %d AND status != 2", time
```

_Ý nghĩa của lệnh query_: Cập nhật status của tất cả transaction thành 4 nếu time của transaction >= time được truyền vào và status của transaction đó khác 2.

> _Output_: return YES ( nếu cập nhật status thành công) or NO ( nếu cập nhật status thất bại hoặc có lỗi xảy ra trong quá trình cập nhật)

## getLastTransaction(address)

Hàm sử dụng để lấy ra transaction cuối cùng, có address là tham số được truyền vào.

> Input:
>
> - address: Địa chỉ của transaction

Các lệnh query:

```
	@"SELECT * FROM transactionTB WHERE address =  \"%@\" ORDER BY id DESC LIMIT 1",address
```

_Ý nghĩa của lệnh query_: Lấy ra tất cả những thông tin của transaction có địa chỉ bằng địa chỉ được truyền vào và chỉ lấy 1 transaction trên cùng duy nhất và sắp xếp giảm dần theo cột id.

> Output: return resultArray(1 mảng mà trong đó chứa transaction duy nhất lấy được thành công) or nil (nếu lấy không thành công hoặc xảy ra lỗi trong quá trình thực hiện)

## getLastTransactionWithStatus(address, status)

Hàm sử dụng để lấy ra transaction có address và status tương ứng với tham số được truyền vào

> Input:
>
> - address: địa chỉ của transaction muốn get
> - status: status của transaction muốn get

Các lệnh query:

```
	@"SELECT * FROM transactionTB WHERE address = \"%@\" AND status = %d ORDER BY id DESC LIMIT 1",address, status
```

_Ý nghĩa của lệnh query_: lấy ra tất của thông tin của transaction có địa chỉ, status bằng với tham số truyền vào và chỉ lấy duy nhất 1 transaction trên cùng duy nhất và sắp xếp giảm dần theo cột id.

> Output: return resultArray(1 mảng mà trong đó chứa transaction duy nhất lấy được thành công) or nil (nếu lấy không thành công hoặc xảy ra lỗi trong quá trình thực hiện)

## insertSmartContract(data)

Hàm sử dụng để thêm 1 smart contract mới với data của smart contract được truyền vào bao gồm: name, address, abiData, binData, image, status.

> Input:
>
> - data: data bao gồm các thông tin của smart contract như: name, address, abiDAta, binData, image.

Các lệnh query:

```
	@"INSERT INTO smartContractTB(name, address, abiData, binData, image, status) values (\"%@\",\"%@\", \"%@\", \"%@\", \"%@\", %d)",data[@"name"], data[@"address"], data[@"abiData"], data[@"binData"], data[@"image"], [status intValue]
```

_Ý nghĩa của lệnh query_: Thêm 1 smart contract mới với các thông tin sau: name, address, abiData, bunData, image, status với các giá trị đó được truyền vào từ data.

> Output: return YES (nếu thêm mới smart contract thành công, hoặc địa chỉ của smart đã tồn tại) or NO (Không thêm được smart contract hoặc xảy ra lỗi trong quá trình thực hiện)

## getAllSmartContracts()

Hàm sử dụng để get tất cả smart contract.

> Input: _No Input_

Các lệnh query:

```
	@"SELECT * FROM d_app_table WHERE type = 2 AND status = 1 AND isInstalled = 1 ORDER BY id ASC"
```

_Ý nghĩa của lệnh query_: Lấy tất cả thông tin của tất cả smart contract có type = 2, status = 1, isInstalled = 1.

> Output: return resultArray( nếu get thành công danh sách smart contract) or @[] (nếu không có smart contract nào trong db)

## getAllWallets()

Hàm sử dụng để lấy ra tất cả các wallet trong db.

> Input: _No Input_

Các lệnh query:

```
	@"SELECT * FROM walletTB"
```

> Output: return resultArray( nếu get thành công danh sách smart contract) or @[] (nếu không có smart contract nào trong db)

## getAllTransactions()

Hàm sử dụng để lấy ra tất cả các transaction có trong db.

> Input: _No Input_

Các lệnh query:

```
	@"SELECT * FROM transactionTB"
```

> Output: return resultArray( nếu get thành công danh sách transaction) or @[] (nếu không có transaction nào trong db)

## getAllWhiteList()

Hàm sử dụng để lấy ra tất cả các whitelist

> Input: _No Input_

Các lệnh query:

```
	@"SELECT * FROM whiteListTB"
```

> Output: return resultArray( nếu get thành công danh sách whitelist) or @[] (nếu không có whitelist nào trong db)

## updateSmartContractStatusByAddress(status, address)

Hàm sử dụng để cập nhật status của smart contract có address tương ứng với address được truyền vào

> Input:
>
> - status: trạng thái muốn thay đổi cho smart contract.
> - address: địa chỉ của smart contract.

Các lệnh query:

```
	@"UPDATE smartContractTB SET status = %d where address = \"%@\"", status, address
```

_Ý nghĩa của lệnh query_: Cập nhật status của smart contract thành status được truyền vào với điều kiện smart contract đó có address là address được truyền vào.

> Output: return YES( nếu update thành công) or NO ( nếu update không thành công hoặc xảy ra lỗi trong quá trình thực hiện)

## getWalletPagination(offset, limit)

Hàm để get danh sách wallet theo offset và limit

> Input:
>
> - offset: vị trí bắt đầu thực hiện query
> - limit: giới hạn số lượng row khi thực hiện query

Các lệnh query:

```
	@"SELECT * FROM walletTB ORDER BY id ASC LIMIT %d OFFSET %ld ", limit, offset
```

_Ý nghĩa của lệnh query_: Get danh sách wallet và sắp xếp theo id tăng dần, thực hiện get từ offset được truyền vào và giới hạn get cũng từ giá trị limit được truyền vào từ bên ngoài.

> Output: return resultArray( nếu get thành công danh sách wallet) or @[] (nếu không có wallet nào trong db)

## countWalletTable()

Hàm sử dụng đếm số lượng wallet trong bảng wallet.

> Input: _No Input_

Các lệnh query:

```
	@"SELECT COUNT(*) FROM walletTB"
```

_Ý nghĩa của lệnh query_: Đếm số lượng row trong bảng wallet.

> Output: return count( số lượng wallet trong bảng wallet)

## getSmartContractPagination(offset, limit)

Hàm để get danh sách smart contract theo offset và limit

> Input:
>
> - offset: vị trí bắt đầu thực hiện query
> - limit: giới hạn số lượng row khi thực hiện query

Các lệnh query:

```
	@"SELECT * FROM smartContractTB  where status = 1 ORDER BY id ASC LIMIT %d OFFSET %ld ", limit, offset
```

_Ý nghĩa của lệnh query_: Get danh sách smart contract nào có trạng thái là 1 và sắp xếp theo id tăng dần, thực hiện get từ offset được truyền vào và giới hạn get cũng từ giá trị limit được truyền vào từ bên ngoài.

> Output: return resultArray( nếu get thành công danh sách smart contract) or @[] (nếu không có smart contract nào trong db)

## countSmartContractTable()

Hàm sử dụng đếm số lượng smart contract trong bảng smartContract.

> Input: _No Input_

Các lệnh query:

```
	@"SELECT COUNT(*) FROM smartContractTB"
```

_Ý nghĩa của lệnh query_: Đếm số lượng row trong bảng smartContract.

> Output: return count( số lượng smart contract trong bảng smartContract)

## getTransactionPagination(offset, limit,address)

Hàm để get danh sách smart contract theo offset và limit

> Input:
>
> - offset: vị trí bắt đầu thực hiện query
> - limit: giới hạn số lượng row khi thực hiện query
> - address: địa chỉ của transaction muốn get

Các lệnh query:

```
	@"SELECT * FROM transactionTB WHERE (address= \"%@\" AND type = \"send\") OR (toAddress= \"%@\" AND type = \"receive\") ORDER BY id DESC LIMIT %d OFFSET %ld ", address, address, limit, offset
```

_Ý nghĩa của lệnh query_: Get danh sách transaction với điều kiện là address của transaction đó bằng với address được truyền vào và type = send hoặc có toAddress bằng address được truyền vào và type = receive, với bảng transaction và sắp xếp theo thứ tự giảm dần của id và thực hiện get từ offset được truyền vào và giới hạn get cũng từ giá trị limit được truyền vào từ bên ngoài.

> Output: return resultArray( nếu get thành công danh sách transaction) or @[] (nếu không có transaction nào thỏa mãn điều kiện query)

## countTransactionTable()

Hàm sử dụng đếm số lượng transaction trong transaction.

> Input: _No Input_

Các lệnh query:

```
	@"SELECT COUNT(*) FROM transactionTB"
```

_Ý nghĩa của lệnh query_: Đếm số lượng row trong bảng transaction.

> Output: return count( số lượng transaction trong bảng transaction)

## getWhiteListPagination(offset, limit)

Hàm để get danh sách whitelist theo offset và limit

> Input:
>
> - offset: vị trí bắt đầu thực hiện query
> - limit: giới hạn số lượng row khi thực hiện query

Các lệnh query:

```
	@"SELECT * FROM whiteListTB ORDER BY id ASC LIMIT %d OFFSET %ld ", limit, offset
```

_Ý nghĩa của lệnh query_: Get danh sách whitelist và sắp xếp theo id tăng dần, thực hiện get từ offset được truyền vào và giới hạn get cũng từ giá trị limit được truyền vào từ bên ngoài.

> Output: return resultArray( nếu get thành công danh sách whitelist) or @[] (nếu không có whitelist nào trong db)

## insertWhiteList(data)

Hàm sử dụng để thêm 1 whitelist mới với data của whitelist được truyền vào bao gồm: image, name, email, user_name, phoneNumber, address.

> Input:
>
> - data: data bao gồm các thông tin của whitelist như: image, name, email, user_name, phoneNumber, address.

Các lệnh query:

```
	@"INSERT INTO whiteListTB(image, name, email, user_name, phoneNumber, address) values (\"%@\",\"%@\", \"%@\", \"%@\", \"%@\", \"%@\")", data[@"image"], data[@"name"], data[@"email"], data[@"user_name"], data[@"phoneNumber"], data[@"address"]
```

_Ý nghĩa của lệnh query_: Thêm 1 whitelist mới vào bảng whitelist với các thông tin sau: image, name, email, user_name, phoneNumber, address với các giá trị đó được truyền vào từ data.

> Output: return YES (nếu thêm mới whitelist thành công, hoặc địa chỉ của whitelist đã tồn tại) or NO (Không thêm được whitelist hoặc xảy ra lỗi trong quá trình thực hiện)

## countWhiteListTable()

Hàm sử dụng đếm số lượng whitelist trong db.

> Input: _No Input_

Các lệnh query:

```
	@"SELECT COUNT(*) FROM whiteListTB"
```

_Ý nghĩa của lệnh query_: Đếm số lượng row trong bảng whitelist.

> Output: return count( số lượng whitelist trong bảng whitelist).

## getTransactionByHash(hash)

Hàm sử dụng để cập nhật lại trạng thái của 1 transaction dựa vào

> Input:
>
> - hash: Mã hash của transaction muốn get

Các lệnh query:

```
   @"SELECT * FROM transactionTB WHERE hash= \"%@\" ORDER BY id DESC", hash]
```

_Ý nghĩa của lệnh query_: Get transacion có mã hash bằng với mã hash được truyền vào và sắp xếp theo thứ tự giảm dần.

> _Output_: return [resultArray objectAtIndex:0] ( transaction được get thành công) or nil(nếu không get được transaction hoặc xảy ra lỗi trong quá trình get);

## checkWhiteList(address)

Hàm sử dụng để kiểm tra xem whitelist có address tương ứng với address được truyền vào có tồn tại hay không

> Input:
>
> - address: address của whitelist muốn get

Các lệnh query:

```
   @"SELECT * FROM whiteListTB WHERE address= \"%@\" ORDER BY id DESC", address
```

_Ý nghĩa của lệnh query_: Get whitlist có address bằng với address được truyền vào và sắp xếp theo thứ tự giảm dần của id.

> _Output_: return YES( tồn tại whitelist có chứa address mong muốn) or NO(nếu không tìm thấy whitelist nào có chứa address mong muốn hoặc có lỗi xảy ra trong quá trình get);

## getTransactionPaginationWithStatus(offset, limit, address, status)

Hàm sử dụng để get danh sách transaction có address, status tương ứng với tham số truyền vào và get theo offset, limit.

> Input:
>
> - offset: vị trí bắt đầu thực hiện query
> - limit: giới hạn số lượng row khi thực hiện query
> - address: address của transaction muốn get
> - status: status của transaction muốn get

Các lệnh query:

```
   @"SELECT * FROM transactionTB WHERE (address= \"%@\" AND type = \"send\" AND status = %d) OR (toAddress= \"%@\" AND type = \"receive\" AND status = %d) ORDER BY id DESC LIMIT %d OFFSET %ld ", address, status, address, status, limit, offset
```

_Ý nghĩa của lệnh query_: Get những transaction nào có address bằng với address được truyền vào, type = send và status bằng status được truyền vào hoặc toAddress bằng với address được truyền vào, tpe = receive và status bằng status được truyền vào. Và sắp xếp theo thứ tự giảm dần của id, thực hiện get từ offset được truyền vào và giới hạn get cũng từ giá trị limit được truyền vào từ bên ngoài.

> _Output_: return resultArray( danh sách transaction thỏa mãn điều kiện query) or @[](nếu không tìm thấy transaction nào thỏa mãn điều kiện query hoăc có lỗi xảy ra trong quá trình thực hiện);

## countTransactionTableWithStatus(status)

Hàm sử dụng để đếm số lượng transaction có status tương ứng với status được truyền vào

> Input:
>
> - status: status của transaction muốn đếm

Các lệnh query:

```
   @"SELECT COUNT(*) FROM transactionTB WHERE status = %d", status
```

_Ý nghĩa của lệnh query_: Đếm xem có bao nhiêu transaction có status bằng với status được truyền vào.

> _Output_: return count( số lượng transaction có status bằng với status được truyền vào);

## insertRecentNode(data)

Hàm sử dụng để thêm 1 recent node mới với data được truyền vào bao gồm: ip, port, time.

> Input:
>
> - data: data của recent node trong đó bao gồm các thông tin: ip, port, time.

Các lệnh query:

```
   @"INSERT INTO recentNodeTB(ip, port, time) values (\"%@\", %ld, %ld)",
	data[@"ip"], [data[@"port"] longValue], [data[@"time"] longValue]
```

_Ý nghĩa của lệnh query_: Thêm 1 recent node vào bảng recentNode với value của các field ip, port, time được truyền vào từ data.

> _Output_: return YES( thêm thành công recent node vào bảng recentNode) or NO( thêm không thành công recent node hoặc có lỗi xảy ra trong quá trình thực hiện);

## updateRecentNode(data)

<!-- Hàm sử dụng để cập nhật 1 recent node với data được truyền vào bao gồm: ip, port, time.

> Input:
>
> - data: data của recent node trong đó bao gồm các thông tin: ip, port, time.

Các lệnh query:

```
   @"UPDATE recentNodeTB SET time = %ld WHERE ip = \"%@\" AND port = %ld", [data[@"time"] integerValue], data[@"ip"], [data[@"port"] integerValue]
```

_Ý nghĩa của lệnh query_: Thêm 1 recent node vào bảng recentNode với value của các field ip, port, time được truyền vào từ data.

> _Output_: return YES( thêm thành công recent node vào bảng recentNode) or NO( thêm không thành công recent node hoặc có lỗi xảy ra trong quá trình thực hiện); -->

## checkRecentNodeExist(ip, port)

Hàm sử dụng để kiểm tra node có tồn tại hay không.

> Input:
>
> - ip: ip của recent node.
> - port: port của recent node.

Các lệnh query:

```
   @"SELECT * FROM recentNodeTB WHERE ip = \"%@\" AND port = %d", ip, port
```

_Ý nghĩa của lệnh query_: Lấy ra node có ip và port bằng với ip và port truyền vào.

> _Output_: return YES( nếu tồn tại node thỏa mãn điều kiện query) or NO( nếu không có node thỏa mãn điều kiện query hoặc có lỗi xảy ra trong quá trình thực hiện)

## getNodeRecent()

Hàm sử dụng để get tất cả các recent node

> Input: _No Input_

Các lệnh query:

```
	@"SELECT * FROM recentNodeTB ORDER BY time ASC LIMIT 10 OFFSET 0"
```

_Ý nghĩa của lệnh query_: get danh sách recent node và sắp xếp theo thứ tự time tăng dần, giới hạn phần tử là 10 và lấy bắt đầu từ vị trí đầu tiên của table.

> Output: return resultArray( danh sách 10 node đầu tiên) or @[] ( nếu không có phần node nào trong bảng recentNode hoặc xảy ra lỗi trong quá trình thực hiện).

## deleteWalletByAddress(address)

Hàm sử dụng xóa wallet với address tương ứng với address truyền vào.

> Input:
>
> - address: address của wallet muốn xóa.

Các lệnh query:

```
   @"DELETE FROM walletTB WHERE address = \"%@\"", address
```

_Ý nghĩa của lệnh query_: Xóa wallet có address bằng với address truyền vào.

> _Output_: return YES( nếu xóa thành công wallet) or NO( nếu xóa không thành công hoặc có lỗi xảy ra trong quá trình thực hiện)

## deleteTransactionByFromAddress(address)

Hàm sử dụng xóa transaction với address tương ứng với address truyền vào.

> Input:
>
> - address: address của transaction muốn xóa.

Các lệnh query:

```
   @"DELETE FROM transactionTB WHERE address = \"%@\"", address
```

_Ý nghĩa của lệnh query_: Xóa transaction có address bằng với address truyền vào.

> _Output_: return YES( nếu xóa thành công transaction) or NO( nếu xóa không thành công hoặc có lỗi xảy ra trong quá trình thực hiện)

## getDAppByBundleId(bundleId)

Hàm sử dụng get app với bundleId tương ứng với bundleId truyền vào.

> Input:
>
> - bundleId: bundleId của app muốn get.

Các lệnh query:

```
   @"SELECT * FROM decentralizedApplicationTB WHERE bundleId= \"%@\" ORDER BY id DESC", bundleId
```

_Ý nghĩa của lệnh query_: get app có bundleId bằng với bundleId được truyền vào.

> _Output_: return [resultArray objectAtIndex:0]( nếu get app thành công) or nil( nếu get app không thành công hoặc có lỗi xảy ra trong quá trình thực hiện)

## getDAppBrowserById(dAppId)

Hàm sử dụng get browser app với dAppId tương ứng với dAppId truyền vào.

> Input:
>
> - dAppId: dAppId của app muốn get.

Các lệnh query:

```
   @"SELECT * FROM decentralizedApplicationTB WHERE id= %d AND type = 1 AND isInstalled = 1", dAppId
```

_Ý nghĩa của lệnh query_: get app có dAppId bằng với dAppId được truyền vào.

> _Output_: return [resultArray objectAtIndex:0]( nếu get app thành công) or nil( nếu get app không thành công hoặc có lỗi xảy ra trong quá trình thực hiện)

## getAllDAppByGroupId(groupId)

Hàm sử dụng get tất cả DApp với groupId tương ứng với groupId truyền vào.

> Input:
>
> - groupId: groupId của những app muốn get.

Các lệnh query:

```
   @"SELECT * FROM decentralizedApplicationTB WHERE type = 1 AND isInstalled = 1 AND isShowInApp = 1 AND groupId = %d ORDER BY page, position", groupId
```

_Ý nghĩa của lệnh query_: get những app có dAppId type = 1, isInstalled = 1, isShowInApp = 1 , groupId bằng groupId được truyền vào và sắp xếp theo page và position.

> _Output_: return resultArray( danh sách app thoả mãn điều kiện query) or @[]( nếu không có app nào thỏa mãn điều kiện query hoặc có lỗi xảy ra trong quá trình thực hiện)

## insertDApp(data)

Hàm sử dụng để thêm 1 DApp mới.

> Input:
>
> - data: data của DApp muốn thêm mới bao gồm: name, author, hash, sign, version, logo, pathStorage, time, totalWallet, totalTransaction, size, bundleId, orientation, urlRoot, urlLoadingScreen, urlLauchScreen, isInjectJs, urlWeb, isLocal, fullScreen, statusBar, groupId, isShowInApp, page, position, isInstalled, abiData, binData, status, type.

Các lệnh query:

```
   @"INSERT INTO decentralizedApplicationTB(name, author, hash, sign, version, logo, pathStorage, time, totalWallet, totalTransaction, size, bundleId, orientation, urlRoot, urlLoadingScreen, urlLauchScreen, isInjectJs, urlWeb, isLocal, fullScreen, statusBar, groupId, isShowInApp, page, position, isInstalled, abiData, binData, status, type) values (\"%@\",\"%@\", \"%@\", \"%@\", \"%@\", \"%@\", \"%@\", %ld, %ld, %ld, \"%@\", \"%@\", \"%@\", \"%@\", \"%@\", \"%@\", %ld, \"%@\", %ld, %ld, \"%@\", %ld, %ld, %ld, %ld, %ld, \"%@\", \"%@\", %ld, %ld)",
	data[@"name"], data[@"author"], data[@"hash"], data[@"sign"], data[@"version"], data[@"logo"], data[@"pathStorage"], [data[@"time"] integerValue], [data[@"totalWallet"] integerValue], [data[@"totalTransaction"] integerValue], data[@"size"], data[@"id"], data[@"orientation"], data[@"urlRoot"] != NULL ? data[@"urlRoot"] : @"", data[@"urlLoadingScreen"], data[@"urlLauchScreen"], [data[@"isInjectJs"] integerValue], data[@"urlWeb"], [data[@"isLocal"] integerValue], [data[@"fullScreen"] integerValue], data[@"statusBar"], [data[@"groupId"] integerValue], [data[@"isShowInApp"] integerValue], [data[@"page"] integerValue], [data[@"position"] integerValue], [data[@"isInstalled"] integerValue], data[@"abiData"], data[@"abiData"], [data[@"status"] integerValue], [data[@"type"] integerValue]
```

_Ý nghĩa của lệnh query_: Thêm 1 DApp mới với data được truyền vào.

> _Output_: return
>
> - Thông báo đã tồn tại (Nếu DApp đó đã tồn tại)
> - Thông báo success (Nếu thêm DApp thành công)
> - Thông báo fail nếu thêm DApp không thành công
> - nil nếu (Xảy ra lỗi trong quá trình thêm DApp)

## updateDApp(data)

Hàm sử dụng để cập nhật thông tin cho DApp.

> Input:
>
> - data: data của DApp muốn thêm mới bao gồm: name, author, hash, sign, version, logo, pathStorage, time, totalWallet, totalTransaction, size, bundleId, orientation, urlRoot, urlLoadingScreen, urlLauchScreen, isInjectJs, urlWeb, isLocal, fullScreen, statusBar, groupId, isShowInApp, page, position, isInstalled, abiData, binData, status, type.

Các lệnh query:

```
   @"UPDATE decentralizedApplicationTB SET name = \"%@\", author = \"%@\", hash = \"%@\", sign = \"%@\", version = \"%@\", logo = \"%@\", pathStorage = \"%@\", time = %ld, totalWallet = %ld, totalTransaction = %ld, size = %@, orientation = \"%@\", urlRoot = \"%@\", urlLoadingScreen = \"%@\", urlLauchScreen = \"%@\", groupId = %ld WHERE bundleId = \"%@\"",
    data[@"name"], data[@"author"], data[@"hash"], data[@"sign"], data[@"version"], data[@"logo"], data[@"pathStorage"], [data[@"time"] integerValue], [data[@"totalWallet"] integerValue], [data[@"totalTransaction"] integerValue], data[@"size"], data[@"orientation"], data[@"urlRoot"], data[@"urlLoadingScreen"], data[@"urlLauchScreen"], [data[@"groupId"] integerValue], data[@"id"]
```

_Ý nghĩa của lệnh query_: Cập nhật thông tin của DApp với data được truyền vào.

> _Output_: return
>
> - Thông báo success (Nếu udpate DApp thành công)
> - Thông báo fail ( Nếu udpate DApp không thành công)
> - nil nếu (Xảy ra lỗi trong quá trình update DApp)

## updateDAppGroupId(groupId,dAppId)

Hàm sử dụng để cập nhật groupId cho DApp.

> Input:
>
> - groupId: id của group muốn thay đổi DApp

Các lệnh query:

```
   @"UPDATE decentralizedApplicationTB SET groupId = %ld WHERE id = %d ", [groupId integerValue], dAppId
```

_Ý nghĩa của lệnh query_: Cập nhật group id của DApp có id tương ứng với id được truyền vào.

> _Output_: return
>
> - Thông báo success (Nếu udpate DApp thành công)
> - Thông báo fail ( Nếu udpate DApp không thành công)
> - nil nếu (Xảy ra lỗi trong quá trình update DApp)

## updateDAppPage(page, dAppId)

Hàm sử dụng để cập nhật page của DApp.

> Input:
>
> - page: page muốn thay đổi cho DApp
> - dAppId: id của DApp muốn thay đổi

Các lệnh query:

```
  @"UPDATE decentralizedApplicationTB SET page = %d WHERE id = %d ", [page intValue], dAppId]
```

_Ý nghĩa của lệnh query_: Cập nhật page của DApp có id tương ứng với id được truyền vào.

> _Output_: return
>
> - Thông báo success (Nếu udpate DApp thành công)
> - Thông báo fail ( Nếu udpate DApp không thành công)
> - nil nếu (Xảy ra lỗi trong quá trình update DApp)

## updateDAppPosition(position, dAppId)

Hàm sử dụng để cập nhật vị trí cho DApp.

> Input:
>
> - position: vị trí muốn thay đổi của DApp
> - dAppId: id của DApp muốn thay đổi

Các lệnh query:

```
   @"UPDATE decentralizedApplicationTB SET position = %d WHERE id = %d ", [position intValue], dAppId]
```

_Ý nghĩa của lệnh query_: Cập nhật group id của DApp có id tương ứng với id được truyền vào.

> _Output_: return
>
> - Thông báo success (Nếu udpate DApp thành công)
> - Thông báo fail ( Nếu udpate DApp không thành công)
> - nil nếu (Xảy ra lỗi trong quá trình update DApp)

## updateDAppStatusByBundleId(status, bundleId)

Hàm sử dụng để cập nhật trạng thái cho DApp.

> Input:
>
> - status: trạng thái muốn thay đổi cho DApp
> - bundleId: bundleId của DApp muốn thay đổi

Các lệnh query:

```
   @"UPDATE decentralizedApplicationTB SET status = %d WHERE type = 2 AND bundleId = \"%@\" ", [status intValue], bundleId]
```

_Ý nghĩa của lệnh query_: Cập nhật trạng thái của DApp có cps type = 2 và bundleId tương ứng với bundleId được truyền vào.

> _Output_: return
>
> - Thông báo success (Nếu udpate DApp thành công)
> - Thông báo fail ( Nếu udpate DApp không thành công)
> - nil nếu (Xảy ra lỗi trong quá trình update DApp)

## getDAppPagination(offset, limit)

Hàm sử dụng để get danh sách DApp theo offset và limit

> Input:
>
> - offset: vị trí bắt đầu thực hiện query
> - limit: giới hạn số lượng row khi thực hiện query

Các lệnh query:

```
	@"SELECT * FROM decentralizedApplicationTB WHERE type = 1 AND isInstalled = 1 AND isShowInApp = 1 ORDER BY page, position LIMIT %d OFFSET %ld ", limit, offset]
```

_Ý nghĩa của lệnh query_: Get danh sách DApp với type = 1, isInstalled = 1, isShowInApp = 1 và sắp xếp theo page tăng dần, thực hiện get từ offset được truyền vào và giới hạn get cũng từ giá trị limit được truyền vào từ bên ngoài.

> Output: return resultArray( nếu get thành công danh sách DApp) or @[] (nếu không có DApp nào thỏa mãn điều kiện query hoặc có lỗi xảy ra trong quá trình thực hiện)

## getAllDapps()

Hàm sử dụng để get tất cả DApp

> Input: _No Input_

Các lệnh query:

```
	@"SELECT * FROM decentralizedApplicationTB WHERE type = 1 AND isInstalled = 1 AND isShowInApp = 1 ORDER BY page, position"
```

_Ý nghĩa của lệnh query_: Lấy tất cả các DApp có type = 1, isInstalled = 1, isShowInApp = 1 và sắp xếp theo thứ tự page, position tăng dần.

> Output: return resultArray( nếu get thành công danh sách DApp) or @[] (nếu không có DApp nào thỏa mãn điều kiện query hoặc xảy ra lỗi trong quá trình thực hiện)

## checkDAppExist(bundleId)

Hàm sử dụng để kiểm tra DApp có tồn tại hay không.

> Input:
>
> - bundleId: bundleId của DApp cần kiểm tra.

Các lệnh query:

```
   @"SELECT * FROM decentralizedApplicationTB WHERE bundleId= \"%@\" ORDER BY id DESC", bundleId]
```

_Ý nghĩa của lệnh query_: Lấy ra DApp có bundleId bằng với bundleId được truyền vào và sắp xếp theo tứ tự giảm dần của id.

> _Output_: return YES( nếu tồn tại DApp thỏa mãn điều kiện query) or NO( nếu không có DApp thỏa mãn điều kiện query hoặc có lỗi xảy ra trong quá trình thực hiện)

## deleteDApp(idApp)

Hàm sử dụng để xóa DApp có id tương ứng với id được truyền vào.

> Input:
>
> - idApp: id của DApp cần xóa.

Các lệnh query:

```
   @"DELETE FROM decentralizedApplicationTB WHERE id = %d", idApp]
```

_Ý nghĩa của lệnh query_: Xóa DApp có id bằng với id được truyền vào.

> _Output_: return YES( nếu xóa DApp thành công) or NO( nếu xóa DApp thất bại hoặc có lỗi xảy ra trong quá trình thực hiện)

## getMaxPositionDApp()

Hàm sử dụng để get vị trí lớn nhất trong danh sách DApp.

> Input: _No Input_

Các lệnh query:

```
   @"SELECT position FROM decentralizedApplicationTB ORDER BY position DESC LIMIT 1"
```

_Ý nghĩa của lệnh query_: Lấy cột tất cả position của tất cả DApp và sắp xếp theo thứ tự giảm dần sau đó lấy ra position đầu tiên( position lớn nhất).

> _Output_: return count( Vị trí lớn nhất của DApp)

## countDappTable()

Hàm sử dụng để đếm tổng số lượng App.

> Input: _No Input_

Các lệnh query:

```
   @"SELECT COUNT(*) FROM decentralizedApplicationTB"]
```

_Ý nghĩa của lệnh query_: Đếm số lượng DApp có trong bảng decentralizedApplicationTB.

> _Output_: return count( Số lượng DApp)

## deleteAllDApp()

Hàm sử dụng để xóa tất cả các App.

> Input: _No Input_

Các lệnh query:

```
   @"DELETE FROM decentralizedApplicationTB"
```

_Ý nghĩa của lệnh query_: Xóa tất cả DApp có trong bảng decentralizedApplicationTB.

> _Output_: return YES( Nếu xóa tất cả DApp thành công) or NO(Nếu xóa tất cả các DApp thất bại hoặc có lỗi xảy ra trong quá trình thực hiện)

## insertGroup(data)

Hàm sử dụng để thêm group với data được truyền vào.

> Input:
>
> - data: data của group cần thêm vào bao gồm: name, position.

Các lệnh query:

```
   @"INSERT INTO GroupDApp(name, position) values (\"%@\", %ld)"
```

_Ý nghĩa của lệnh query_: Thêm group mới vào bảng GroupDApp với name, postion được truyền vào từ data.

> _Output_: return YES( nếu thêm mới group thành công) or NO( Nếu thêm group mới thất bại hoặc xảy ra lỗi trong quá trình thưc hiện)

## deleteGroupById(groupId)

Hàm sử dụng để xóa group với group id được truyền vào.

> Input:
>
> - groupId: groupId của group cần xóa.

Các lệnh query:

```
   @"DELETE FROM GroupDApp WHERE id = %d", groupId]
```

_Ý nghĩa của lệnh query_: Xóa group có id bằng với id được truyền vào.

> _Output_: return YES( nếu xóa group thành công) or NO( Nếu xóa group thất bại hoặc xảy ra lỗi trong quá trình thưc hiện)

## getAllGroupDApps()

Hàm sử dụng để get tất cả các group.

> Input: _No Input_

Các lệnh query:

```
   @"SELECT * FROM GroupDApp"
```

_Ý nghĩa của lệnh query_: Lấy tất cả group có trong bảng GroupDApp.

> _Output_: return resultArray( danh sách group) or @[](Nếu gt group thất bại hoặc có lỗi xảy ra trong quá trình thực hiện)

## renameGroupById(name, groupId)

Hàm sử dụng để đổi tên group với name và groupId được truyền vào để thực hiện query.

> Input:
>
> - name: name muốn đổi cho group.
> - groupId: id của group cần đổi tên.

Các lệnh query:

```
   @"UPDATE GroupDApp SET name = \"%@\" WHERE id = %d", name, [groupId intValue]
```

_Ý nghĩa của lệnh query_: Cập nhật tên của group(có id bằng với id được truyền vào) bằng tên được truyền vào .

> _Output_: return YES( nếu cập nhật tên group thành công) or NO( nếu cập nhật tên group thất bại hoặc có lỗi xảy ra trong quá trình thực hiện)

## updateGroupPosition(position, groupId)

Hàm sử dụng để cập nhật vị trí của group với vị trí và group id được truyền vào để thực hiện query.

> Input:
>
> - position: position muốn đổi cho group.
> - groupId: id của group cần đổi tên.

Các lệnh query:

```
   @"UPDATE GroupDApp SET position = %d WHERE id = %d ", [position intValue], groupId
```

_Ý nghĩa của lệnh query_: Cập nhật position của group(có id bằng với id được truyền vào) bằng position được truyền vào .

> _Output_: return
>
> - Thông báo success (Nếu udpate vị trí group thành công)
> - Thông báo fail ( Nếu udpate vị trí group không thành công)
> - nil nếu (Xảy ra lỗi trong quá trình update DApp)

## getMaxPositionGroup()

Hàm sử dụng để get vị trí lớn nhất của group.

> Input: _No Input_

Các lệnh query:

```
   @"SELECT position FROM GroupDApp ORDER BY position DESC LIMIT 1"
```

_Ý nghĩa của lệnh query_: Lấy tất cả position của tất cả group trong bảng GroupDApp sắp xếp theo thứ tự giảm dần của position và sau đó lấy ra phần tử đầu tiên (phần tử lớn nhất).

> _Output_: return count( vị trí lớn nhất của group)

## getLastestGroupID()

Hàm sử dụng để get id lớn nhất của group.

> Input: _No Input_

Các lệnh query:

```
   @"SELECT MAX(id) FROM GroupDApp"
```

_Ý nghĩa của lệnh query_: Lấy ra id lớn nhất trong tất cả các group.

> _Output_: return identity( id lớn nhất trong tất cả các group)
