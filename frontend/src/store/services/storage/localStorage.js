import { isJson } from '../../../utils/helpers'


export const loadState = () => {
  try {
    const serializedState = localStorage.getItem('state')
    if (serializedState === null) {
      return undefined;
    }
    return JSON.parse(serializedState)
  } catch (err) {
    return undefined
  }
}

export const saveState = (state) => {
  try {
    const serializedState = JSON.stringify(state)
    localStorage.setItem('state', serializedState)
  } catch (err) {
    // to define
  }
}

export const loadValue = (key) => {
  try {
    const value = localStorage.getItem(key)

    return value
  } catch (err) {
    console.log(err)
  }
}

export const setValue = (key, value) => {
  try {
    localStorage.setItem(key, value)

  } catch (err) {
    console.log(err)
  }
}

export const setShowHelpModalSetting = (value) => {
  setValue('showHelpModal', value)
}

export const loadShowHelpModalSetting = () => {
  let value = loadValue('showHelpModal')

  if (value === null || value === "true") {
    return true
  } else {
    return false
  }
}

