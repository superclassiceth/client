// @flow
import React from 'react'
import {normal as proofNormal} from '../constants/tracker'
import {Box, Button, FollowButton} from '../common-adapters'
import {globalMargins} from '../styles/style-guide'
import type {Props} from './user-actions'

export default function UserActions ({trackerState, currentlyFollowing, style, onFollow, onUnfollow, onAcceptProofs}: Props) {
  if (trackerState === proofNormal) {
    if (currentlyFollowing) {
      return (
        <Box style={style}>
          <FollowButton following onUnfollow={onUnfollow} style={{marginRight: 0}} />
        </Box>
      )
    } else {
      return (
        <Box style={style}>
          <FollowButton following={false} onFollow={onFollow} style={{marginRight: 0}} />
        </Box>
      )
    }
  } else {
    return (
      <Box style={style}>
        <Button type='Unfollow' label='Untrack' onClick={onUnfollow} style={{marginRight: globalMargins.xtiny}} />
        <Button type='Follow' label='Accept' onClick={onAcceptProofs} style={{marginRight: 0}} />
      </Box>
    )
  }
}
