export class ErrorHolder {
  errors = new Map<string, Set<string>>();

  checkIsSet(field: string) {
    return this.errors.has(field);
  }

  addValidation(evaluation: boolean, field: string, msg: string) {
    if (evaluation) {
      this.setError(field, msg);
    } else {
      this.clearError(field, msg);
    }
  }

  setError(field: string, msg: string) {
    let msgs = new Set<string>();

    if (this.checkIsSet(field)) {
      msgs = this.errors.get(field)!;
    } else {
      console.log(`Campo "${field}" no está inicializado`);
    }

    if (msgs.has(msg)) {
      console.log(`Campo "${field}" ya contiene "${msg}"`)
      return;
    }

    this.errors.set(field, msgs.add(msg));
  }

  clearError(field: string, msg: string) {
    if (!this.checkIsSet(field)) {
      console.log(`Campo "${field}" no está inicializado`);
      return;
    }

    const msgs = this.errors.get(field)!;
    const elementRemoved = msgs.delete(msg);
    if (!elementRemoved) {
      console.log(`Campo "${field}" no contiene mensaje "${field}"`);
    }

    if (msgs.size == 0) {
      this.errors.delete(field);
      return;
    }

    this.errors.set(field, msgs);
  }
  
  getErrorsList() {
    const errorsList = [];
    for (const [field, msgs] of this.errors.entries()) {
      errorsList.push({ field, msgs: Array.from(msgs) });
    }
    return errorsList;
  }

  displayErrors() {
    console.log(this.errors);
  }

  clone() {
    const newHolder = new ErrorHolder();
    newHolder.errors = this.errors;
    return newHolder;
  }

  passedValidations() {
    for (const [_, msgs] of this.errors.entries()) {
      if (msgs.size > 0) {
        return false;
      }
    }
    return true;
  }

  hasErrors() {
    return this.errors.size > 0;
  }
}

