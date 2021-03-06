// Copyright 2020 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


import { computed, observable } from 'mobx';
import { fromPromise, IPromiseBasedObservable } from 'mobx-utils';

import { AppState } from '../../context/app_state_provider';
import { consumeContext, provideContext } from '../../libs/context';
import * as iter from '../../libs/iter_utils';
import { streamTestBatches, streamTestExonerationBatches, streamTestResultBatches, streamVariantBatches, TestLoader } from '../../models/test_loader';
import { TestNode, VariantStatus } from '../../models/test_node';
import { Expectancy, Invocation } from '../../services/resultdb';

/**
 * Records state of the invocation page.
 */
export class InvocationPageState {
  @observable.ref appState!: AppState;
  @observable.ref invocationId = '';
  @observable.ref selectedTabId = '';

  @computed
  get invocationName(): string {
    return 'invocations/' + this.invocationId;
  }

  @computed
  get invocationReq(): IPromiseBasedObservable<Invocation> {
    if (!this.appState.resultDb) {
      // Returns a promise that never resolves when resultDb isn't ready.
      return fromPromise(new Promise(() => {}));
    }
    return fromPromise(this.appState.resultDb.getInvocation({name: this.invocationName}));
  }

  @computed
  get invocation(): Invocation | null {
    if (this.invocationReq.state !== 'fulfilled') {
      return null;
    }
    return this.invocationReq.value;
  }

  @observable.ref selectedNode!: TestNode;
  @observable.ref showExpected = false;
  @observable.ref showExonerated = true;
  @observable.ref showFlaky = true;

  @computed private get testResultBatchIterFn() {
    if (!this.appState?.resultDb) {
      return async function*() {};
    }
    return iter.teeAsync(streamTestResultBatches(
      {
        invocations: [this.invocationName],
        predicate: {
          expectancy: this.showExpected ? Expectancy.All : Expectancy.VariantsWithUnexpectedResults,
        },
        readMask: '*',
      },
      this.appState.resultDb,
    ));
  }

  @computed private get testExonerationBatchIterFn() {
    if (!this.appState?.resultDb) {
      return async function*() {};
    }
    return iter.teeAsync(streamTestExonerationBatches(
      {invocations: [this.invocationName]},
      this.appState.resultDb,
    ));
  }

  @computed
  private get testIterFn() {
    let variantBatches = streamVariantBatches(
      this.testResultBatchIterFn(),
      this.testExonerationBatchIterFn(),
    );

    variantBatches = this.showExonerated ?
      variantBatches :
      iter.mapAsync(variantBatches, (batch) => batch.filter((v) => v.status !== VariantStatus.Exonerated));

    // Known Issue:
    // A variant's status may change after filtering from expected/unexpected to
    // flaky if a result with a different expected value is received in the next
    // batch. In that case, some flaky variants are not filtered out.
    // This should be a rare occurrence. Since usually, results of the same test
    // variant should be in the same batch.
    variantBatches = this.showFlaky ?
      variantBatches :
      iter.mapAsync(variantBatches, (batch) => batch.filter((v) => v.status !== VariantStatus.Flaky));

    return iter.teeAsync(streamTestBatches(variantBatches));
  }

  @computed({keepAlive: true})
  get testLoader() { return new TestLoader(TestNode.newRoot(), this.testIterFn()); }
}

export const consumePageState = consumeContext<'pageState', InvocationPageState>('pageState');
export const providePageState = provideContext<'pageState', InvocationPageState>('pageState');
