/*
 * Copyright 2018 The CovenantSQL Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package limits defines limits of the CovenantSQL system.
package limits

const (
	// MaxTxBroadcastTTL defines the TTL limit of a AddTx request broadcasting within the
	// block producers.
	MaxTxBroadcastTTL = 1
	// MaxPendingTxsPerAccount defines the limit of pending transactions of one account.
	MaxPendingTxsPerAccount = 1000
)
