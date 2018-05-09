import "testing"

func BenchmarkWithdrawals(b *testing.B) {
	// ...

	server := NewFundServer(b.N)

	// ...

	// Spawn off the workers
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < dollarsPerFounder; i++ {
				server.Commands <- WithdrawCommand{Amount: 1}
			}
		}()
	}

	// ...

	balanceResponseChan := make(chan int)
	server.Commands <- BalanceCommand{Response: balanceResponseChan}
	balance := <-balanceResponseChan

	if balance != 0 {
		b.Error("Balance wasn't zero:", balance)
	}
}